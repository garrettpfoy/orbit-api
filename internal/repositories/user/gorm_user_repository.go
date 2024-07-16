package user

import (
	"garrettpfoy/orbit-api/internal/models"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func (r *GormUserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepository) GetUser(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *GormUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *GormUserRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
