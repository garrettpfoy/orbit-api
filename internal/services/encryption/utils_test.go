package encryption_test

import (
	"garrettpfoy/orbit-api/internal/services/encryption"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEncryptionService(t *testing.T) {
	secret := "abcdefghijklmnopqrstuvwxyz123456" // 32 bytes for AES-256
	service := encryption.NewEncryptionService(secret)

	assert.NotNil(t, service, "Service should not be nil")
}

func TestEncryptDecrypt(t *testing.T) {
	secret := "abcdefghijklmnopqrstuvwxyz123456" // 32 bytes for AES-256
	service := encryption.NewEncryptionService(secret)

	originalText := "Hello, World!"
	encryptedText, err := service.Encrypt(originalText)
	assert.NoError(t, err, "Encrypt should not return an error")

	decryptedText, err := service.Decrypt(encryptedText)
	assert.NoError(t, err, "Decrypt should not return an error")
	assert.Equal(t, originalText, decryptedText, "Decrypted text should match the original")
}

func TestEncryptWithShortSecret(t *testing.T) {
	secret := "shortsecret" // Less than 16/24/32 bytes
	service := encryption.NewEncryptionService(secret)

	_, err := service.Encrypt("Hello, World!")
	assert.Error(t, err, "Encrypt should return an error with a short secret")
}

func TestDecryptInvalidHex(t *testing.T) {
	secret := "abcdefghijklmnopqrstuvwxyz123456" // 32 bytes for AES-256
	service := encryption.NewEncryptionService(secret)

	_, err := service.Decrypt("invalidhex")
	assert.Error(t, err, "Decrypt should return an error with invalid hex input")
}
