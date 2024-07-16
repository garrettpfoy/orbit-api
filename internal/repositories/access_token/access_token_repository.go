package access_token

import (
	"garrettpfoy/orbit-api/internal/models"
)

type AccessTokenRepository interface {
	CreateAccessToken(token *models.AccessToken) error
	GetAccessToken(id uint) (*models.AccessToken, error)
	GetAccessTokenByUserID(userID uint) (*models.AccessToken, error)
	UpdateAccessToken(token *models.AccessToken) error
	DeleteAccessToken(id uint) error
}
