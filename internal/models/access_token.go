package models

import (
	"garrettpfoy/orbit-api/internal/services/encryption"
	"time"

	"gorm.io/gorm"
)

type AccessToken struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"not null"`
	User         User      `gorm:"foreignKey:UserID"`
	AccessToken  string    `gorm:"unique;not null"`
	RefreshToken string    `gorm:"unique;not null"`
	ExpiryTime   time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

var encryptionService encryption.OrbitEncryptionService

func SetEncryptionService(service encryption.OrbitEncryptionService) {
	encryptionService = service
}

func (token *AccessToken) BeforeSave(tx *gorm.DB) (err error) {
	if token.AccessToken != "" {
		token.AccessToken, err = encryptionService.Encrypt(token.AccessToken)
		if err != nil {
			return err
		}
	}
	if token.RefreshToken != "" {
		token.RefreshToken, err = encryptionService.Encrypt(token.RefreshToken)
		if err != nil {
			return err
		}
	}
	return nil
}

func (token *AccessToken) AfterFind(tx *gorm.DB) (err error) {
	if token.AccessToken != "" {
		token.AccessToken, err = encryptionService.Decrypt(token.AccessToken)
		if err != nil {
			return err
		}
	}
	if token.RefreshToken != "" {
		token.RefreshToken, err = encryptionService.Decrypt(token.RefreshToken)
		if err != nil {
			return err
		}
	}
	return nil
}
