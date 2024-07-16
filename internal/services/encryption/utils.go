package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// OrbitEncryptionService defines the interface for encryption and decryption
type OrbitEncryptionService interface {
	Encrypt(text string) (string, error)
	Decrypt(cryptoText string) (string, error)
}

// EncryptionService implements the OrbitEncryptionService interface
type EncryptionService struct {
	secret []byte
}

// NewEncryptionService creates a new instance of EncryptionService
func NewEncryptionService(secret string) OrbitEncryptionService {
	return &EncryptionService{
		secret: []byte(secret),
	}
}

// Encrypt encrypts the given text using AES
func (e *EncryptionService) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher(e.secret)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given ciphertext using AES
func (e *EncryptionService) Decrypt(cryptoText string) (string, error) {
	ciphertext, err := hex.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.secret)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
