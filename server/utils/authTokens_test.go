package utils_test

import (
	"task-inator3000/utils"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		isRefresh bool
		wantErr   bool
	}{
		{"AuthToken", "tester@testing.com", false, false},
		{"RefreshToken", "tester@testing.com", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := utils.CreateJWT(tt.email, tt.isRefresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJWT() returned error: %v", err)
			}
		})
	}
}

func TestVerifyJWT(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		isRefresh bool
		wantErr   bool
	}{
		{"AuthToken", "tester@testing.com", false, false},
		{"RefreshToken", "tester@testing.com", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := utils.CreateJWT(tt.email, tt.isRefresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJWT() returned error: %v", err)
			}

			result, err := utils.VerifyJWT(token, tt.isRefresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyJWT() returned error: %v", err)

				if result != tt.email {
					t.Errorf("Expected: %v, Received: %v", tt.email, result)
				}
			}
		})
	}
}
