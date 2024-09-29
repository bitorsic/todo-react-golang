package utils_test

import (
	"os"
	"task-inator3000/utils"
	"testing"
)

const plaintext string = "This is my plaintext for encryption"

func TestAESEncrypt(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		wantErr bool
	}{
		{"Normal", "TheEncKeyForAESMustBe32LongOkay?", false},
		{"ShortKey", "ThisIsNot32Long", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("AES_KEY", tt.key)

			_, err := utils.AESEncrypt(plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESEncrypt() returned error: %v", err)
			}
		})
	}
}

func TestAESDecrypt(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		wantErr bool
	}{
		{"Normal", "TheEncKeyForAESMustBe32LongOkay?", false},
		{"ShortKey", "ThisIsNot32Long", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("AES_KEY", tt.key)

			ciphertext, err := utils.AESEncrypt(plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESEncrypt() returned error: %v", err)
			}

			result, err := utils.AESDecrypt(ciphertext)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESDecrypt() returned error: %v", err)

				if plaintext != result {
					t.Errorf("Expected: %v, Received: %v", plaintext, result)
				}
			}
		})
	}
}
