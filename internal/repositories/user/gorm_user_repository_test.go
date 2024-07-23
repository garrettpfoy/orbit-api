package user_test

import (
	"errors"
	"garrettpfoy/orbit-api/internal/models"
	"garrettpfoy/orbit-api/internal/repositories/user"
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

	err = db.AutoMigrate(&models.User{}, &models.Session{}, &models.AccessToken{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := user.NewGormUserRepository(db)

	user := &models.User{
		SpotifyUserID: newString("spotify123"),
		Email:         newString("user@example.com"),
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)

	var createdUser models.User
	err = db.First(&createdUser, user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, *user.SpotifyUserID, *createdUser.SpotifyUserID)
	assert.Equal(t, *user.Email, *createdUser.Email)
}

func TestGetUserByID(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := user.NewGormUserRepository(db)

	user := &models.User{
		SpotifyUserID: newString("spotify123"),
		Email:         newString("user@example.com"),
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, *user.SpotifyUserID, *retrievedUser.SpotifyUserID)
	assert.Equal(t, *user.Email, *retrievedUser.Email)
}

func TestGetUserBySpotifyID(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := user.NewGormUserRepository(db)

	user := &models.User{
		SpotifyUserID: newString("spotify123"),
		Email:         newString("user@example.com"),
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetUserBySpotifyID(*user.SpotifyUserID)
	assert.NoError(t, err)
	assert.Equal(t, *user.SpotifyUserID, *retrievedUser.SpotifyUserID)
	assert.Equal(t, *user.Email, *retrievedUser.Email)
}

func TestGetUserByEmail(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := user.NewGormUserRepository(db)

	user := &models.User{
		SpotifyUserID: newString("spotify123"),
		Email:         newString("user@example.com"),
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetUserByEmail(*user.Email)
	assert.NoError(t, err)
	assert.Equal(t, *user.SpotifyUserID, *retrievedUser.SpotifyUserID)
	assert.Equal(t, *user.Email, *retrievedUser.Email)
}

func TestGetUserSessions(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := user.NewGormUserRepository(db)

	session1 := &models.Session{Slug: "slug1", HostID: 1}
	session2 := &models.Session{Slug: "slug2", HostID: 1}

	err = db.Create(session1).Error
	assert.NoError(t, err)
	err = db.Create(session2).Error
	assert.NoError(t, err)

	user := &models.User{
		SpotifyUserID: newString("spotify123"),
		Email:         newString("user@example.com"),
		Sessions:      []*models.Session{session1, session2},
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)

	sessions, err := repo.GetUserSessions(user.ID)
	assert.NoError(t, err)
	assert.Len(t, sessions, 2)
}

func TestUpdateUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := user.NewGormUserRepository(db)

	user := &models.User{
		SpotifyUserID: newString("spotify123"),
		Email:         newString("user@example.com"),
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)

	user.SpotifyUserID = newString("spotify456")
	err = repo.UpdateUser(user)
	assert.NoError(t, err)

	var updatedUser models.User
	err = db.First(&updatedUser, user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, *user.SpotifyUserID, *updatedUser.SpotifyUserID)
}

func TestDeleteUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := user.NewGormUserRepository(db)

	user := &models.User{
		SpotifyUserID: newString("spotify123"),
		Email:         newString("user@example.com"),
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)

	err = repo.DeleteUser(user.ID)
	assert.NoError(t, err)

	var deletedUser models.User
	err = db.First(&deletedUser, user.ID).Error
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func newString(s string) *string {
	return &s
}
