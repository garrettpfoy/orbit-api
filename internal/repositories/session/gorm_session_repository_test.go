package session_test

import (
	"errors"
	"garrettpfoy/orbit-api/internal/models"
	"garrettpfoy/orbit-api/internal/repositories/session"
	"garrettpfoy/orbit-api/internal/services/encryption"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	models.SetEncryptionService(encryption.NewEncryptionService("abcdefghijklmnopqrstuvwxyz123456"))

	err = db.AutoMigrate(&models.Session{}, &models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateSession(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := session.NewGormSessionRepository(db)

	session := &models.Session{
		Slug:   "unique_slug",
		HostID: 1,
		Host:   models.User{Model: gorm.Model{ID: 1}},
	}

	err = repo.CreateSession(session)
	assert.NoError(t, err)

	var createdSession models.Session
	err = db.First(&createdSession, session.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, session.Slug, createdSession.Slug)
}

func TestGetSessions(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := session.NewGormSessionRepository(db)

	session1 := &models.Session{
		Slug:   "unique_slug_1",
		HostID: 1,
		Host:   models.User{Model: gorm.Model{ID: 1}},
	}
	session2 := &models.Session{
		Slug:   "unique_slug_2",
		HostID: 2,
		Host:   models.User{Model: gorm.Model{ID: 2}},
	}

	err = repo.CreateSession(session1)
	assert.NoError(t, err)
	err = repo.CreateSession(session2)
	assert.NoError(t, err)

	sessions, err := repo.GetSessions()
	assert.NoError(t, err)
	assert.Len(t, sessions, 2)
}

func TestGetSession(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := session.NewGormSessionRepository(db)

	session := &models.Session{
		Slug:   "unique_slug",
		HostID: 1,
		Host:   models.User{Model: gorm.Model{ID: 1}},
	}

	err = repo.CreateSession(session)
	assert.NoError(t, err)

	retrievedSession, err := repo.GetSession(session.ID)
	assert.NoError(t, err)
	assert.Equal(t, session.Slug, retrievedSession.Slug)
}

func TestGetSessionByHostID(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := session.NewGormSessionRepository(db)

	session := &models.Session{
		Slug:   "unique_slug",
		HostID: 1,
		Host:   models.User{Model: gorm.Model{ID: 1}},
	}

	err = repo.CreateSession(session)
	assert.NoError(t, err)

	retrievedSession, err := repo.GetSessionByHostID(session.HostID)
	assert.NoError(t, err)
	assert.Equal(t, session.Slug, retrievedSession.Slug)
}

func TestGetUsersInSession(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := session.NewGormSessionRepository(db)

	user1 := &models.User{Model: gorm.Model{ID: 1}}
	user2 := &models.User{Model: gorm.Model{ID: 2}}

	session := &models.Session{
		Slug:   "unique_slug",
		HostID: 1,
		Host:   models.User{Model: gorm.Model{ID: 1}},
		Users:  []*models.User{user1, user2},
	}

	err = repo.CreateSession(session)
	assert.NoError(t, err)

	users, err := repo.GetUsersInSession(session.ID)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestGetSessionBySlug(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := session.NewGormSessionRepository(db)

	session := &models.Session{
		Slug:   "unique_slug",
		HostID: 1,
		Host:   models.User{Model: gorm.Model{ID: 1}},
	}

	err = repo.CreateSession(session)
	assert.NoError(t, err)

	retrievedSession, err := repo.GetSessionBySlug(session.Slug)
	assert.NoError(t, err)
	assert.Equal(t, session.Slug, retrievedSession.Slug)
}

func TestUpdateSession(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := session.NewGormSessionRepository(db)

	session := &models.Session{
		Slug:   "unique_slug",
		HostID: 1,
		Host:   models.User{Model: gorm.Model{ID: 1}},
	}

	err = repo.CreateSession(session)
	assert.NoError(t, err)

	session.Slug = "updated_slug"
	err = repo.UpdateSession(session)
	assert.NoError(t, err)

	var updatedSession models.Session
	err = db.First(&updatedSession, session.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, session.Slug, updatedSession.Slug)
}

func TestDeleteSession(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := session.NewGormSessionRepository(db)

	session := &models.Session{
		Slug:   "unique_slug",
		HostID: 1,
		Host:   models.User{Model: gorm.Model{ID: 1}},
	}

	err = repo.CreateSession(session)
	assert.NoError(t, err)

	err = repo.DeleteSession(session.ID)
	assert.NoError(t, err)

	var deletedSession models.Session
	err = db.First(&deletedSession, session.ID).Error
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
