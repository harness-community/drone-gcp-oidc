package env

import (
	"fmt"
	"os"
)

func VerifyEnv() error {
	oidcIdToken := os.Getenv("PLUGIN_OIDC_TOKEN_ID")
	projectId := os.Getenv("PLUGIN_PROJECT_ID")
	poolId := os.Getenv("PLUGIN_POOL_ID")
	providerId := os.Getenv("PLUGIN_PROVIDER_ID")
	serviceAccountEmail := os.Getenv("PLUGIN_SERVICE_ACCOUNT_EMAIL_ID")

	if oidcIdToken == "" || projectId == "" || poolId == "" || providerId == "" || serviceAccountEmail == "" {
		return fmt.Errorf("missing required environment variables")
	}

	return nil
}
