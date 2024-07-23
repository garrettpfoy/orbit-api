package models

import "garrettpfoy/orbit-api/internal/services/encryption"

var encryptionService encryption.OrbitEncryptionService

func SetEncryptionService(service encryption.OrbitEncryptionService) {
	encryptionService = service
}
