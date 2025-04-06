package encryption

import (
	"testing"
	"fmt"
	"axolot/src/encryption"
)

func TestEncryptionDecryption(t *testing.T) {
	// Generate the key from host details
	_, err := encryption.GenerateKeyFromHostDetails()
	if err != nil {
		t.Fatalf("Error generating key: %v", err)
	}

	tests := []struct {
		plaintext string
	}{
		{"Username"},
		{"Password"},
		{"A3!)(*K9@#@%#1F5"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Encrypt/Decrypt: %s", tt.plaintext), func(t *testing.T) {
			// Encrypt the plaintext
			ciphertext, err := encryption.EncryptData(tt.plaintext)
			if err != nil {
				t.Errorf("Error encrypting data: %v", err)
			}

			// Decrypt the ciphertext
			decrypted, err := encryption.DecryptData(ciphertext)
			if err != nil {
				t.Errorf("Error decrypting data: %v", err)
			}

			// Verify that decrypted data matches the original plaintext
			if decrypted != tt.plaintext {
				t.Errorf("Decrypted data does not match original plaintext. Got: %s, Want: %s", decrypted, tt.plaintext)
			}
		})
	}
}
