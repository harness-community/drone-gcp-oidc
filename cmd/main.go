package main

import (
	"drone/plugin/gcp-oidc/internal/gcp"
	"fmt"
	"os"
)

func main() {
	oidcIdToken := os.Getenv("PLUGIN_OIDC_ID_TOKEN")
	projectId := os.Getenv("PLUGIN_PROJECT_NUMBER")
	poolId := os.Getenv("PLUGIN_POOL_ID")
	providerId := os.Getenv("PLUGIN_PROVIDER_ID")
	serviceAccountEmail := os.Getenv("PLUGIN_SERVICE_ACCOUNT_EMAIL")

	if oidcIdToken == "" || projectId == "" || poolId == "" || providerId == "" || serviceAccountEmail == "" {
		fmt.Println("Missing required environment variables")
		os.Exit(1)
	}

	federalToken, err := gcp.GetFederalToken(oidcIdToken, projectId, poolId, providerId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	accessToken, err := gcp.GetGoogleCloudAccessToken(federalToken, serviceAccountEmail)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully retrieved access token")
	fmt.Println(accessToken)

	os.Setenv("GOOGLE_ACCESS_TOKEN", accessToken)

	outputFile, err := os.OpenFile(os.Getenv("DRONE_OUTPUT"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	_, err = fmt.Fprintf(outputFile, "GOOGLE_ACCESS_TOKEN=%s\n", accessToken)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully wrote access token to environment variables")
}
