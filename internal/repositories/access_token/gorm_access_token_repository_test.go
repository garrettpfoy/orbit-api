package access_token_test

import (
	"errors"
	"garrettpfoy/orbit-api/internal/models"
	repository "garrettpfoy/orbit-api/internal/repositories/access_token"
	"garrettpfoy/orbit-api/internal/services/encryption"
	"testing"
	"time"

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

	err = db.AutoMigrate(&models.AccessToken{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateAccessToken(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := repository.NewGormAccessTokenRepository(db)

	token := &models.AccessToken{
		UserID:       1,
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiryTime:   time.Now().Add(time.Hour),
		SessionID:    1,
	}

	err = repo.CreateAccessToken(token)
	assert.NoError(t, err)

	createdToken, err := repo.GetAccessToken(token.ID)
	assert.NoError(t, err)

	assert.Equal(t, "access_token", createdToken.AccessToken)
}

func TestGetAccessToken(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := repository.NewGormAccessTokenRepository(db)

	token := &models.AccessToken{
		UserID:       1,
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiryTime:   time.Now().Add(time.Hour),
		SessionID:    1,
	}

	err = repo.CreateAccessToken(token)
	assert.NoError(t, err)

	retrievedToken, err := repo.GetAccessToken(token.ID)
	assert.NoError(t, err)
	assert.Equal(t, "access_token", retrievedToken.AccessToken)
}

func TestGetAccessTokenByUserID(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := repository.NewGormAccessTokenRepository(db)

	token := &models.AccessToken{
		UserID:       1,
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiryTime:   time.Now().Add(time.Hour),
		SessionID:    1,
	}

	err = repo.CreateAccessToken(token)
	assert.NoError(t, err)

	retrievedToken, err := repo.GetAccessTokenByUserID(token.UserID)
	assert.NoError(t, err)
	assert.Equal(t, "access_token", retrievedToken.AccessToken)
}

func TestGetAccessTokenBySessionID(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := repository.NewGormAccessTokenRepository(db)

	token := &models.AccessToken{
		UserID:       1,
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiryTime:   time.Now().Add(time.Hour),
		SessionID:    1,
	}

	err = repo.CreateAccessToken(token)
	assert.NoError(t, err)

	retrievedToken, err := repo.GetAccessTokenBySessionID(token.SessionID)
	assert.NoError(t, err)
	assert.Equal(t, "access_token", retrievedToken.AccessToken)
}

func TestUpdateAccessToken(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := repository.NewGormAccessTokenRepository(db)

	token := &models.AccessToken{
		UserID:       1,
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiryTime:   time.Now().Add(time.Hour),
		SessionID:    1,
	}

	err = repo.CreateAccessToken(token)
	assert.NoError(t, err)

	token.AccessToken = "new_access_token"
	err = repo.UpdateAccessToken(token)
	assert.NoError(t, err)

	var updatedToken models.AccessToken
	err = db.First(&updatedToken, token.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "new_access_token", updatedToken.AccessToken)
}

func TestDeleteAccessToken(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := repository.NewGormAccessTokenRepository(db)

	token := &models.AccessToken{
		UserID:       1,
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiryTime:   time.Now().Add(time.Hour),
		SessionID:    1,
	}

	err = repo.CreateAccessToken(token)
	assert.NoError(t, err)

	err = repo.DeleteAccessToken(token.ID)
	assert.NoError(t, err)

	var deletedToken models.AccessToken
	err = db.First(&deletedToken, token.ID).Error
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
