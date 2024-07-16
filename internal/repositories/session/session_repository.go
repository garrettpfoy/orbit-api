package session

import (
	"garrettpfoy/orbit-api/internal/models"
)

type SessionRepository interface {
	CreateSession(session *models.Session) error
	UpsertSession(session *models.Session) error
	GetSessions() ([]models.Session, error)
	GetSession(id uint) (*models.Session, error)
	UpdateSession(session *models.Session) error
	DeleteSession(id uint) error
}
