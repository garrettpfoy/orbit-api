package models

import (
	"time"

	"gorm.io/gorm"
)

type AccessToken struct {
	gorm.Model
	// User ID represents the user that the token belongs to
	UserID uint `gorm:"not null"`
	// User represents the user that the token belongs to, derived from UserID
	User User
	// AccessToken is the encrypted access token that is used to make requests to the Spotify API
	AccessToken string `gorm:"unique;not null"`
	// RefreshToken is the encrypted refresh token that is used to refresh the access token
	RefreshToken string `gorm:"unique;not null"`
	// ExpiryTime is the time that the access token expires
	ExpiryTime time.Time `gorm:"not null"`
	// Session ID represents the session that the token belongs to (one-to-one relationship)
	SessionID uint `gorm:"not null"`
	// Session represents the session that the token belongs to, derived from SessionID
	Session Session
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
