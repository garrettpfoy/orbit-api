package models

import "gorm.io/gorm"

// User represents the users table
type User struct {
	gorm.Model
	Email         string `gorm:"unique;not null"`
	OAuthProvider string `gorm:"not null"`
}
