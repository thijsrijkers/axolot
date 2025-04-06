package credentialstore

import (
	"testing"
	"axolot/src/encryption"
	"axolot/src/store"
)

func TestCredentialStore(t *testing.T) {
	// Generate the key for encryption/decryption
	_, err := encryption.GenerateKeyFromHostDetails()
	if err != nil {
		t.Fatalf("Error generating key: %v", err)
	}

	// Create a new CredentialStore instance
	cs := credentialstore.New()

	// Test case for adding and verifying credentials
	t.Run("Add and Verify Valid Credentials", func(t *testing.T) {
		username := "testuser"
		password := "testpassword"

		// Add the credentials
		err := cs.AddCredential(username, password)
		if err != nil {
			t.Errorf("Error adding credentials: %v", err)
		}

		// Verify the credentials
		valid := cs.VerifyCredential(username, password)
		if !valid {
			t.Errorf("Expected credentials to be valid for user '%s', but they were not", username)
		}
	})

	// Test case for invalid credentials (wrong username)
	t.Run("Verify Invalid Username", func(t *testing.T) {
		username := "wronguser"
		password := "testpassword"

		// Add valid credentials first
		err := cs.AddCredential("testuser", password)
		if err != nil {
			t.Errorf("Error adding credentials: %v", err)
		}

		// Verify with wrong username
		valid := cs.VerifyCredential(username, password)
		if valid {
			t.Errorf("Expected credentials to be invalid for user '%s', but they were valid", username)
		}
	})

	// Test case for invalid credentials (wrong password)
	t.Run("Verify Invalid Password", func(t *testing.T) {
		username := "testuser"
		wrongPassword := "wrongpassword"

		// Add valid credentials first
		err := cs.AddCredential(username, "testpassword")
		if err != nil {
			t.Errorf("Error adding credentials: %v", err)
		}

		// Verify with wrong password
		valid := cs.VerifyCredential(username, wrongPassword)
		if valid {
			t.Errorf("Expected credentials to be invalid for user '%s' with wrong password, but they were valid", username)
		}
	})

	// Test case for edge case of empty username or password
	t.Run("Verify Empty Username or Password", func(t *testing.T) {
		// Add empty username and password
		err := cs.AddCredential("", "")
		if err != nil {
			t.Errorf("Error adding empty credentials: %v", err)
		}

		// Verify with empty username and password
		valid := cs.VerifyCredential("", "")
		if !valid {
			t.Errorf("Expected credentials to be valid for empty username and password, but they were not")
		}

		// Verify with non-empty username and empty password
		valid = cs.VerifyCredential("nonemptyusername", "")
		if valid {
			t.Errorf("Expected credentials to be invalid for empty password, but they were valid")
		}

		// Verify with empty username and non-empty password
		valid = cs.VerifyCredential("", "nonemptypassword")
		if valid {
			t.Errorf("Expected credentials to be invalid for empty username, but they were valid")
		}
	})
}
