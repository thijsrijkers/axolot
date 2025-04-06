package credentialstore

import (
    "axolot/src/encryption"
	"fmt"
)

type CredentialStore struct {
	store map[string][]byte
}

func New() *CredentialStore {
	return &CredentialStore{
		store: make(map[string][]byte),
	}
}

func (cs *CredentialStore) AddCredential(username, password string) error {
	// Encrypt the username
	encryptedUsername, err := encryption.EncryptData(username)
	if err != nil {
		return fmt.Errorf("failed to encrypt username: %v", err)
	}

	// Encrypt the password
	encryptedPassword, err := encryption.EncryptData(password)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %v", err)
	}

	// Convert encrypted password from string to []byte (base64 encoded)
	encryptedPasswordBytes := []byte(encryptedPassword)

	// Store the encrypted username as the key and encrypted password as the value
	cs.store[encryptedUsername] = encryptedPasswordBytes
	return nil
}

func (cs *CredentialStore) VerifyCredential(username, password string) bool {
	// Encrypt the input username to compare with stored encrypted usernames
	encryptedUsername, err := encryption.EncryptData(username)
	if err != nil {
		return false
	}

	// Retrieve the encrypted password associated with the encrypted username
	encryptedPassword, exists := cs.store[encryptedUsername]
	if !exists {
		return false
	}

	// Decrypt the password (encryptedPassword is []byte, so we need to decode it to string for decryption)
	encryptedPasswordString := string(encryptedPassword)

	// Decrypt the password to compare with the input password
	decryptedPassword, err := encryption.DecryptData(encryptedPasswordString)
	if err != nil {
		return false
	}

	// Verify that the decrypted password matches the input password
	return decryptedPassword == password
}