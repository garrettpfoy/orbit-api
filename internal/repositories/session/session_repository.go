package session

import (
	"garrettpfoy/orbit-api/internal/models"
)

type SessionRepository interface {
	// CreateSession validates a session and creates a new session in the database
	CreateSession(session *models.Session) error
	// GetSessions retrieves all sessions from the database
	GetSessions() ([]models.Session, error)
	// GetSession retrieves a session from the database by its ID
	GetSession(id uint) (*models.Session, error)
	// GetSessionByUserID retrieves a session from the database by its user ID, if it exists
	GetSessionByHostID(userID uint) (*models.Session, error)
	// GetUsersInSession retrieves all users in a session by the session ID
	GetUsersInSession(sessionID uint) ([]models.User, error)
	// GetSessionBySlug retrieves a session from the database by its slug
	UpdateSession(session *models.Session) error
	// DeleteSession deletes a session from the database by its ID
	DeleteSession(id uint) error
}
