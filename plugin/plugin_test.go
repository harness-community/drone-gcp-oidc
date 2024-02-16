// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import "testing"

func TestVerifyEnv(t *testing.T) {
	type args struct {
		Args Args
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "missing oidc-token",
			args: args{
				Args: Args{
					OIDCToken:  "",
					ProjectID:  "project-id",
					PoolID:     "pool-id",
					ProviderID: "provider-id",
					ServiceAcc: "service-account",
				}},
			wantErr: true,
		},
		{
			name: "missing project-id",
			args: args{
				Args: Args{
					OIDCToken:  "oidc-token",
					ProjectID:  "",
					PoolID:     "pool-id",
					ProviderID: "provider-id",
					ServiceAcc: "service-account",
				}},
			wantErr: true,
		},
		{
			name: "missing pool-id",
			args: args{
				Args: Args{
					OIDCToken:  "oidc-token",
					ProjectID:  "project-id",
					PoolID:     "",
					ProviderID: "provider-id",
					ServiceAcc: "service-account",
				}},
			wantErr: true,
		},
		{
			name: "missing provider-id",
			args: args{
				Args: Args{
					OIDCToken:  "oidc-token",
					ProjectID:  "project-id",
					PoolID:     "pool-id",
					ProviderID: "",
					ServiceAcc: "service-account",
				}},
			wantErr: true,
		},
		{
			name: "missing service-account",
			args: args{
				Args: Args{
					OIDCToken:  "oidc-token",
					ProjectID:  "project-id",
					PoolID:     "pool-id",
					ProviderID: "provider-id",
					ServiceAcc: "",
				}},
			wantErr: true,
		},
		{
			name: "all args provided",
			args: args{
				Args: Args{
					OIDCToken:  "oidc-token",
					ProjectID:  "project-id",
					PoolID:     "pool-id",
					ProviderID: "provider-id",
					ServiceAcc: "service-account",
				}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := VerifyEnv(tt.args.Args)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteEnvToFile(t *testing.T) {
	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "write to file",
			args: args{
				key:   "key",
				value: "value",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteEnvToFile(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteEnvToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFederalToken(t *testing.T) {
	type args struct {
		idToken    string
		projectID  string
		poolID     string
		providerID string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get federal token",
			args: args{
				idToken:    "id-token",
				projectID:  "project-id",
				poolID:     "pool-id",
				providerID: "provider-id",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetFederalToken(tt.args.idToken, tt.args.projectID, tt.args.poolID, tt.args.providerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFederalToken() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestGetGoogleCloudAccessToken(t *testing.T) {
	type args struct {
		federatedToken      string
		serviceAccountEmail string
		duration            string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get google cloud access token",
			args: args{
				federatedToken:      "federated-token",
				serviceAccountEmail: "service-account-email",
				duration:            "3600",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetGoogleCloudAccessToken(tt.args.federatedToken, tt.args.serviceAccountEmail, tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoogleCloudAccessToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
