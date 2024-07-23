package models

import "gorm.io/gorm"

// Queue represents the queue table
type Queue struct {
	gorm.Model
	// Track URI is derived from the Spotify API and is used to play the track.
	TrackURI string `gorm:"not null"`
	// Session ID represents the session that the queue item belongs to, which is a foreign key to the sessions table.
	SessionID uint `gorm:"not null"`
	// Session represents the session that the queue item belongs to, derived from the SessionID.
	Session Session
	// User ID represents the user that added the queue item, which is a foreign key to the users table.
	UserID uint `gorm:"not null"`
	// User represents the user that added the queue item, derived from the UserID.
	User User
	// Weight represents the upvote/downvote sum (upvotes - downvotes) for the queue item.
	Weight int `gorm:"default:0;not null"`
}
