package access_token

import (
	"fmt"
	"garrettpfoy/orbit-api/internal/models"

	validator "garrettpfoy/orbit-api/internal/services/validation"

	"gorm.io/gorm"
)

type GormAccessTokenRepository struct {
	db *gorm.DB
}

func NewGormAccessTokenRepository(db *gorm.DB) *GormAccessTokenRepository {
	return &GormAccessTokenRepository{db: db}
}

func (r *GormAccessTokenRepository) CreateAccessToken(token *models.AccessToken) error {
	err := validator.ValidateAccessToken(*token)
	if err != nil {
		return fmt.Errorf("error validating access token: %w", err)
	}

	if err := r.db.Create(token).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormAccessTokenRepository) GetAccessToken(id uint) (*models.AccessToken, error) {
	var token models.AccessToken
	if err := r.db.First(&token, id).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *GormAccessTokenRepository) GetAccessTokenByUserID(userID uint) (*models.AccessToken, error) {
	var token models.AccessToken
	if err := r.db.Where("user_id = ?", userID).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *GormAccessTokenRepository) GetAccessTokenBySessionID(sessionID uint) (*models.AccessToken, error) {
	var token models.AccessToken
	if err := r.db.Where("session_id = ?", sessionID).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *GormAccessTokenRepository) UpdateAccessToken(token *models.AccessToken) error {
	if err := r.db.Save(token).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormAccessTokenRepository) DeleteAccessToken(id uint) error {
	if err := r.db.Delete(&models.AccessToken{}, id).Error; err != nil {
		return err
	}
	return nil
}
