package user

import (
	"garrettpfoy/orbit-api/internal/models"
	"garrettpfoy/orbit-api/internal/services/validation"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) CreateUser(user *models.User) error {
	if err := validation.ValidateUser(*user); err != nil {
		return err
	}
	return r.db.Create(user).Error
}

func (r *GormUserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Sessions").Preload("AccessToken").First(&user, id).Error
	return &user, err
}

func (r *GormUserRepository) GetUserBySpotifyID(spotifyID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("spotify_user_id = ?", spotifyID).Preload("Sessions").Preload("AccessToken").First(&user).Error
	return &user, err
}

func (r *GormUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).Preload("Sessions").Preload("AccessToken").First(&user).Error
	return &user, err
}

func (r *GormUserRepository) GetUserSessions(userID uint) ([]*models.Session, error) {
	var user models.User
	err := r.db.Preload("Sessions").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return user.Sessions, nil
}

func (r *GormUserRepository) UpdateUser(user *models.User) error {
	if err := validation.ValidateUser(*user); err != nil {
		return err
	}
	return r.db.Save(user).Error
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
