package session

import (
	"garrettpfoy/orbit-api/internal/models"

	"gorm.io/gorm"
)

type GormSessionRepository struct {
	DB *gorm.DB
}

func (r *GormSessionRepository) CreateSession(session *models.Session) error {
	return r.DB.Create(session).Error
}

func (r *GormSessionRepository) UpsertSession(session *models.Session) error {
	return r.DB.Save(session).Error
}

func (r *GormSessionRepository) GetSessions() ([]models.Session, error) {
	var sessions []models.Session
	err := r.DB.Find(&sessions).Error
	return sessions, err
}

func (r *GormSessionRepository) GetSession(id uint) (*models.Session, error) {
	var session models.Session
	err := r.DB.First(&session, id).Error
	return &session, err
}

func (r *GormSessionRepository) UpdateSession(session *models.Session) error {
	return r.DB.Save(session).Error
}

func (r *GormSessionRepository) DeleteSession(id uint) error {
	return r.DB.Delete(&models.Session{}, id).Error
}
