package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"axolot/src/host"
)

var storedKey []byte

func SetKey(key []byte) {
	storedKey = key
}

func GenerateKeyFromHostDetails() ([]byte, error) {
	// If the key is already generated, return it from memory
	if storedKey != nil {
		return storedKey, nil
	}

	// Get host details
	info, err := host.GetHostDetails()
	if err != nil {
		return nil, fmt.Errorf("failed to get host details: %w", err)
	}

	// Combine host details into a single string
	data := fmt.Sprintf("%s%s%s", info.Hostname, info.OS, info.Architecture)

	// Hash the combined string to generate a fixed-size key
	hash := sha256.New()
	hash.Write([]byte(data))
	storedKey = hash.Sum(nil)

	return storedKey, nil
}

func EncryptData(plaintext string) (string, error) {
	// Check if the key has been generated and stored in memory
	if storedKey == nil {
		return "", fmt.Errorf("key is not generated yet")
	}

	// Create AES cipher block
	block, err := aes.NewCipher(storedKey)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Use GCM mode for encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	// Generate a nonce
	nonce := make([]byte, aesGCM.NonceSize())
	// In a real application, use a unique nonce for each encryption
	copy(nonce, "unique_nonce_123")

	// Encrypt the data
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptData(ciphertextBase64 string) (string, error) {
	// Check if the key has been generated and stored in memory
	if storedKey == nil {
		return "", fmt.Errorf("key is not generated yet")
	}

	// Create AES cipher block
	block, err := aes.NewCipher(storedKey)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Use GCM mode for decryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	// Decode the base64 ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	// Extract the nonce
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt data: %w", err)
	}
	return string(plaintext), nil
}
