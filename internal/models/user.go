package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the users table
type User struct {
	gorm.Model
	// Username is an arbitrary string the user decides to go by, and is used to display the user's name on the frontend.
	Username string
	// SpotifyUserID is used to denote users who have signed in with Spotify, and may (if an access token exists) have the ability to interact with the Spotify API and therefore host.
	SpotifyUserID *string
	// Email is used to denote users who have signed in with Google, and allows for the ability to vote on tracks.
	Email *string
	// LastVoteTime is used to denote the last time a user voted on a track, and is used to prevent spam voting.
	LastVoteTime *time.Time
	// Sessions represents the many-to-many relationship between users and sessions that are not hosts (signed in users).
	Sessions []*Session `gorm:"many2many:session_users"`
	// AccessTokens represents the one-to-one relationship between users and access tokens (a user should only have one stored access token at a time
	// since only the SpotifyAPI token is stored/used).
	AccessTokenID *uint
	// AccessToken represents the access token that the user may have stored.
	AccessToken *AccessToken
}
