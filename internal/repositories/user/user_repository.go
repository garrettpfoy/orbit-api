package user

import (
	"garrettpfoy/orbit-api/internal/models"
)

type UserRepository interface {
	// CreateUser validates a user and creates a new user in the database
	CreateUser(user *models.User) error
	// GetUser retrieves a user from the database by its ID
	GetUserByID(id uint) (*models.User, error)
	// GetUserBySpotifyID retrieves a user from the database by its Spotify ID, if it exists
	GetUserBySpotifyID(spotifyID string) (*models.User, error)
	// GetUserByEmail retrieves a user from the database by its email, if it exists
	GetUserByEmail(email string) (*models.User, error)
	// GetUserSessions retrieves all sessions a user is in by the user ID
	GetUserSessions(userID uint) ([]models.Session, error)
	// UpdateUser validates a user and updates the user in the database
	UpdateUser(user *models.User) error
	// DeleteUser deletes a user from the database by its ID
	DeleteUser(id uint) error
}
