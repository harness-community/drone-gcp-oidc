// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/oauth2"
	"google.golang.org/api/iamcredentials/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/sts/v1"
)

type staticTokenSource struct {
	token *oauth2.Token
}

func (s *staticTokenSource) Token() (*oauth2.Token, error) {
	return s.token, nil
}

func GetFederalToken(idToken, projectNumber, poolId, providerId, scope string) (string, error) {
	ctx := context.Background()
	stsService, err := sts.NewService(ctx, option.WithoutAuthentication())
	if err != nil {
		return "", fmt.Errorf("failed to create sts service: %w", err)
	}
	audience := fmt.Sprintf("//iam.googleapis.com/projects/%s/locations/global/workloadIdentityPools/%s/providers/%s", projectNumber, poolId, providerId)
	tokenRequest := &sts.GoogleIdentityStsV1ExchangeTokenRequest{
		GrantType:          "urn:ietf:params:oauth:grant-type:token-exchange",
		SubjectToken:       idToken,
		Audience:           audience,
		Scope:              scope,
		RequestedTokenType: "urn:ietf:params:oauth:token-type:access_token",
		SubjectTokenType:   "urn:ietf:params:oauth:token-type:id_token",
	}
	tokenResponse, err := stsService.V1.Token(tokenRequest).Do()
	if err != nil {
		return "", fmt.Errorf("failed to exchange token: %w", err)
	}

	return tokenResponse.AccessToken, nil
}

func GetGoogleCloudAccessToken(federatedToken string, serviceAccountEmail string, duration string, scope string) (string, error) {
	ctx := context.Background()
	tokenSource := &staticTokenSource{
		token: &oauth2.Token{AccessToken: federatedToken},
	}
	service, err := iamcredentials.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		return "", fmt.Errorf("failed to create iamcredentials service: %w", err)
	}

	name := "projects/-/serviceAccounts/" + serviceAccountEmail
	// Split comma-separated scopes to support multiple OAuth scopes
	// e.g., "https://www.googleapis.com/auth/cloud-platform,https://www.googleapis.com/auth/androidpublisher"
	rb := &iamcredentials.GenerateAccessTokenRequest{
		Scope:    strings.Split(scope, ","),
		Lifetime: duration,
	}

	resp, err := service.Projects.ServiceAccounts.GenerateAccessToken(name, rb).Do()
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %w", err)
	}

	return resp.AccessToken, nil
}
