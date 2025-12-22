// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	OIDCToken   string `envconfig:"PLUGIN_OIDC_TOKEN_ID"`
	ProjectID   string `envconfig:"PLUGIN_PROJECT_ID"`
	PoolID      string `envconfig:"PLUGIN_POOL_ID"`
	ProviderID  string `envconfig:"PLUGIN_PROVIDER_ID"`
	ServiceAcc  string `envconfig:"PLUGIN_SERVICE_ACCOUNT_EMAIL_ID"`
	Duration    string `envconfig:"PLUGIN_DURATION"`
	Scope       string `envconfig:"PLUGIN_SCOPE"`
	CreateCreds bool   `envconfig:"PLUGIN_CREATE_APPLICATION_CREDENTIALS_FILE"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	if err := VerifyEnv(args); err != nil {
		return err
	}

	if args.Duration == "" {
		args.Duration = "3600s"
	} else {
		args.Duration = args.Duration + "s"
	}

	if args.Scope == "" {
		args.Scope = "https://www.googleapis.com/auth/cloud-platform"
	}

	if args.CreateCreds {
		// Note: The 'scope' setting does not apply in credentials file mode.
		// Google's external_account ADC JSON format does not support embedding scopes.
		// Scopes must be configured in your application code when initializing
		// the Google Cloud client libraries. See: https://google.aip.dev/auth/4117
		logrus.Infof("creating credentials file\n")
		credsPath, err := WriteCredentialsToFile(args.OIDCToken, args.ProjectID, args.PoolID, args.ProviderID, args.ServiceAcc)
		if err != nil {
			return err
		}
		logrus.Infof("credentials file written to %s\n", credsPath)

		if err := WriteOutputToFile("GOOGLE_APPLICATION_CREDENTIALS", credsPath); err != nil {
			return err
		}

		logrus.Infof("credentials file set as GOOGLE_APPLICATION_CREDENTIALS\n")
	} else {
		federalToken, err := GetFederalToken(args.OIDCToken, args.ProjectID, args.PoolID, args.ProviderID, args.Scope)
		if err != nil {
			return err
		}

		accessToken, err := GetGoogleCloudAccessToken(federalToken, args.ServiceAcc, args.Duration, args.Scope)

		if err != nil {
			return err
		}

		logrus.Infof("access token retrieved successfully\n")

		if err := WriteSecretOutputToFile("GCLOUD_ACCESS_TOKEN", accessToken); err != nil {
			return err
		}

		logrus.Infof("access token set as GCLOUD_ACCESS_TOKEN\n")
		logrus.Infof("access token written to env\n")
	}

	return nil
}

func VerifyEnv(args Args) error {
	if args.OIDCToken == "" {
		return fmt.Errorf("oidc-token is not provided")
	}
	if args.ProjectID == "" {
		return fmt.Errorf("project-id is not provided")
	}
	if args.PoolID == "" {
		return fmt.Errorf("pool-id is not provided")
	}
	if args.ProviderID == "" {
		return fmt.Errorf("provider-id is not provided")
	}
	if args.ServiceAcc == "" {
		return fmt.Errorf("service account email is not provided")
	}
	return nil
}

func WriteOutputToFile(key, value string) error {
	outputFile, err := os.OpenFile(os.Getenv("DRONE_OUTPUT"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}

	defer outputFile.Close()

	_, err = fmt.Fprintf(outputFile, "%s=%s\n", key, value)
	if err != nil {
		return fmt.Errorf("failed to write to env: %w", err)
	}

	return nil
}

func WriteSecretOutputToFile(key, value string) error {
	outputFile, err := os.OpenFile(os.Getenv("HARNESS_OUTPUT_SECRET_FILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}

	defer outputFile.Close()

	_, err = fmt.Fprintf(outputFile, "%s=%s\n", key, value)
	if err != nil {
		return fmt.Errorf("failed to write to env: %w", err)
	}

	return nil
}

func WriteCredentialsToFile(idToken, projectNumber, workforcePoolID, providerID, serviceAccountEmail string) (string, error) {
	homeDir := os.Getenv("DRONE_WORKSPACE")

	if homeDir == "" || homeDir == "/" {
		fmt.Print("could not get home directory, using /home/harness as home directory")
		homeDir = "/home/harness"
	}

	idTokenDir := fmt.Sprintf("%s/tmp", homeDir)
	err := os.MkdirAll(idTokenDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create tmp directory: %w", err)
	}

	idTokenPath := fmt.Sprintf("%s/id_token", idTokenDir)
	if err := os.WriteFile(idTokenPath, []byte(idToken), 0644); err != nil {
		return "", fmt.Errorf("failed to write idToken to file: %w", err)
	}

	fmt.Printf("idTokenPath: %s\n", idTokenPath)

	credsDir := fmt.Sprintf("%s/.config/gcloud", homeDir)
	err = os.MkdirAll(credsDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create gcloud directory: %w", err)
	}

	// create application default credentials file at $HOME/.config/gcloud/application_default_credentials.json
	credsPath := fmt.Sprintf("%s/application_default_credentials.json", credsDir)

	data := map[string]interface{}{
		"type":                              "external_account",
		"audience":                          fmt.Sprintf("//iam.googleapis.com/projects/%s/locations/global/workloadIdentityPools/%s/providers/%s", projectNumber, workforcePoolID, providerID),
		"subject_token_type":                "urn:ietf:params:oauth:token-type:id_token",
		"token_url":                         "https://sts.googleapis.com/v1/token",
		"service_account_impersonation_url": fmt.Sprintf("https://iamcredentials.googleapis.com/v1/projects/-/serviceAccounts/%s:generateAccessToken", serviceAccountEmail),
		"credential_source": map[string]string{
			"file": idTokenPath,
		},
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal json data: %w", err)
	}

	err = os.WriteFile(credsPath, jsonData, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write to credentials file: %w", err)
	}

	fmt.Printf("credsPath: %s\n", credsPath)

	return credsPath, nil
}
