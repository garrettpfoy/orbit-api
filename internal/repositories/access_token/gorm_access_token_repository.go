package access_token

import (
	"garrettpfoy/orbit-api/internal/models"

	"gorm.io/gorm"
)

type GormAccessTokenRepository struct {
	DB *gorm.DB
}

func (r *GormAccessTokenRepository) CreateAccessToken(token *models.AccessToken) error {
	return r.DB.Create(token).Error
}

func (r *GormAccessTokenRepository) GetAccessToken(id uint) (*models.AccessToken, error) {
	var token models.AccessToken
	err := r.DB.First(&token, id).Error
	return &token, err
}

func (r *GormAccessTokenRepository) GetAccessTokenByUserID(userID uint) (*models.AccessToken, error) {
	var token models.AccessToken
	err := r.DB.Where("user_id = ?", userID).First(&token).Error
	return &token, err
}

func (r *GormAccessTokenRepository) UpdateAccessToken(token *models.AccessToken) error {
	return r.DB.Save(token).Error
}

func (r *GormAccessTokenRepository) DeleteAccessToken(id uint) error {
	return r.DB.Delete(&models.AccessToken{}, id).Error
}
