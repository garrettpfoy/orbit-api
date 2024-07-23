package access_token

import (
	"garrettpfoy/orbit-api/internal/models"
)

// Note: Access token encryption and decryption is handled in the models package
// and is outside the scope of the repository interface implementation.

type AccessTokenRepository interface {
	// Create access token validates the access token and creates it in the database
	CreateAccessToken(token *models.AccessToken) error
	// Get access token retrieves an access token by its ID
	GetAccessToken(id uint) (*models.AccessToken, error)
	// Get access token by user ID retrieves an access token by the user ID
	GetAccessTokenByUserID(userID uint) (*models.AccessToken, error)
	// Get access token by session ID retrieves an access token by the session ID
	GetAccessTokenBySessionID(sessionID uint) (*models.AccessToken, error)
	// Update access token updates an access token in the database
	UpdateAccessToken(token *models.AccessToken) error
	// Delete access token deletes an access token from the database
	DeleteAccessToken(id uint) error
}
