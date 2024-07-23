package session

import (
	"garrettpfoy/orbit-api/internal/models"
	"garrettpfoy/orbit-api/internal/services/validation"

	"gorm.io/gorm"
)

type GormSessionRepository struct {
	db *gorm.DB
}

func NewGormSessionRepository(db *gorm.DB) *GormSessionRepository {
	return &GormSessionRepository{db: db}
}

func (r *GormSessionRepository) CreateSession(session *models.Session) error {
	if err := validation.ValidateSession(*session); err != nil {
		return err
	}

	return r.db.Create(session).Error
}

func (r *GormSessionRepository) GetSessions() ([]models.Session, error) {
	var sessions []models.Session
	err := r.db.Preload("Host").Preload("Users").Find(&sessions).Error
	return sessions, err
}

func (r *GormSessionRepository) GetSession(id uint) (*models.Session, error) {
	var session models.Session
	err := r.db.Preload("Host").Preload("Users").First(&session, id).Error
	return &session, err
}

func (r *GormSessionRepository) GetSessionByHostID(userID uint) (*models.Session, error) {
	var session models.Session
	err := r.db.Where("host_id = ?", userID).Preload("Host").Preload("Users").First(&session).Error
	return &session, err
}

func (r *GormSessionRepository) GetUsersInSession(sessionID uint) ([]*models.User, error) {
	var session models.Session
	err := r.db.Preload("Users").First(&session, sessionID).Error
	return session.Users, err
}

func (r *GormSessionRepository) GetSessionBySlug(slug string) (*models.Session, error) {
	var session models.Session
	err := r.db.Where("slug = ?", slug).Preload("Host").Preload("Users").First(&session).Error
	return &session, err
}

func (r *GormSessionRepository) UpdateSession(session *models.Session) error {
	if err := validation.ValidateSession(*session); err != nil {
		return err
	}

	return r.db.Save(session).Error
}

func (r *GormSessionRepository) DeleteSession(id uint) error {
	return r.db.Delete(&models.Session{}, id).Error
}
