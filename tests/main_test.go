package main

import (
	"drone/plugin/gcp-oidc/env"
	"os"
	"testing"
)

func TestMain_EnvVarsNotSet(t *testing.T) {
	// Save current environment variables
	oidcIdToken := os.Getenv("PLUGIN_OIDC_TOKEN_ID")
	projectId := os.Getenv("PLUGIN_PROJECT_ID")
	poolId := os.Getenv("PLUGIN_POOL_ID")
	providerId := os.Getenv("PLUGIN_PROVIDER_ID")
	serviceAccountEmail := os.Getenv("PLUGIN_SERVICE_ACCOUNT_EMAIL_ID")

	// Clear environment variables
	os.Setenv("PLUGIN_OIDC_TOKEN_ID", "")
	os.Setenv("PLUGIN_PROJECT_ID", "")
	os.Setenv("PLUGIN_POOL_ID", "")
	os.Setenv("PLUGIN_PROVIDER_ID", "")
	os.Setenv("PLUGIN_SERVICE_ACCOUNT_EMAIL_ID", "")

	defer func() {
		// Restore original environment variables
		os.Setenv("PLUGIN_OIDC_TOKEN_ID", oidcIdToken)
		os.Setenv("PLUGIN_PROJECT_ID", projectId)
		os.Setenv("PLUGIN_POOL_ID", poolId)
		os.Setenv("PLUGIN_PROVIDER_ID", providerId)
		os.Setenv("PLUGIN_SERVICE_ACCOUNT_EMAIL_ID", serviceAccountEmail)
	}()

	err := env.VerifyEnv()
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	expectedErrorMessage := "missing required environment variables"
	if got := err.Error(); got != expectedErrorMessage {
		t.Errorf("Expected error message %q, but got %q", expectedErrorMessage, got)
	}
}

func TestMain_EnvVarsSet(t *testing.T) {
	// Save current environment variables
	oidcIdToken := os.Getenv("PLUGIN_OIDC_TOKEN_ID")
	projectId := os.Getenv("PLUGIN_PROJECT_ID")
	poolId := os.Getenv("PLUGIN_POOL_ID")
	providerId := os.Getenv("PLUGIN_PROVIDER_ID")
	serviceAccountEmail := os.Getenv("PLUGIN_SERVICE_ACCOUNT_EMAIL_ID")

	// Clear environment variables
	os.Setenv("PLUGIN_OIDC_TOKEN_ID", "test")
	os.Setenv("PLUGIN_PROJECT_ID", "id")
	os.Setenv("PLUGIN_POOL_ID", "id2")
	os.Setenv("PLUGIN_PROVIDER_ID", "pla")
	os.Setenv("PLUGIN_SERVICE_ACCOUNT_EMAIL_ID", "email")

	defer func() {
		// Restore original environment variables
		os.Setenv("PLUGIN_OIDC_TOKEN_ID", oidcIdToken)
		os.Setenv("PLUGIN_PROJECT_ID", projectId)
		os.Setenv("PLUGIN_POOL_ID", poolId)
		os.Setenv("PLUGIN_PROVIDER_ID", providerId)
		os.Setenv("PLUGIN_SERVICE_ACCOUNT_EMAIL_ID", serviceAccountEmail)
	}()

	err := env.VerifyEnv()
	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}
}
