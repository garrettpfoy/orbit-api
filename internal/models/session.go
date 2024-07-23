package models

import (
	"gorm.io/gorm"
)

// Session represents the sessions table
type Session struct {
	gorm.Model
	// Session Slug is the unique identifier for a session, used in the URL to share with others.
	Slug string `gorm:"unique;not null"`
	// Host ID represents the user that created the session, which is a foreign key to the users table.
	HostID uint `gorm:"not null"`
	// Host represents the user that created the session, derived from the HostID.
	Host User `gorm:"foreignKey:HostID"`
	// Users represents the many-to-many relationship between users and sessions that are not hosts (signed in users).
	Users []*User `gorm:"many2many:session_users"`
}
