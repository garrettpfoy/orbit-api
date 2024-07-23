package queue_test

import (
	"errors"
	"garrettpfoy/orbit-api/internal/models"
	"garrettpfoy/orbit-api/internal/repositories/queue"
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

	err = db.AutoMigrate(&models.Queue{}, &models.Session{}, &models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateQueueItem(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := queue.NewGormQueueRepository(db)

	queueItem := &models.Queue{
		TrackURI:  "spotify:track:123",
		SessionID: 1,
		UserID:    1,
		Weight:    10,
	}

	err = repo.CreateQueueItem(queueItem)
	assert.NoError(t, err)

	var createdQueueItem models.Queue
	err = db.First(&createdQueueItem, queueItem.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, queueItem.TrackURI, createdQueueItem.TrackURI)
	assert.Equal(t, queueItem.SessionID, createdQueueItem.SessionID)
	assert.Equal(t, queueItem.UserID, createdQueueItem.UserID)
	assert.Equal(t, queueItem.Weight, createdQueueItem.Weight)
}

func TestGetQueueItemsBySessionID(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := queue.NewGormQueueRepository(db)

	queueItem1 := &models.Queue{
		TrackURI:  "spotify:track:123",
		SessionID: 1,
		UserID:    1,
		Weight:    10,
	}
	queueItem2 := &models.Queue{
		TrackURI:  "spotify:track:456",
		SessionID: 1,
		UserID:    2,
		Weight:    20,
	}

	err = repo.CreateQueueItem(queueItem1)
	assert.NoError(t, err)
	err = repo.CreateQueueItem(queueItem2)
	assert.NoError(t, err)

	prioritize := true
	queueItems, err := repo.GetQueueItemsBySessionID(1, &prioritize)
	assert.NoError(t, err)
	assert.Len(t, queueItems, 2)
	assert.Equal(t, queueItem2.TrackURI, queueItems[0].TrackURI) // queueItem2 should come first due to higher weight
	assert.Equal(t, queueItem1.TrackURI, queueItems[1].TrackURI)
}

func TestGetQueueItemsByUserID(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := queue.NewGormQueueRepository(db)

	queueItem1 := &models.Queue{
		TrackURI:  "spotify:track:123",
		SessionID: 1,
		UserID:    1,
		Weight:    10,
	}
	queueItem2 := &models.Queue{
		TrackURI:  "spotify:track:456",
		SessionID: 2,
		UserID:    1,
		Weight:    20,
	}

	err = repo.CreateQueueItem(queueItem1)
	assert.NoError(t, err)
	err = repo.CreateQueueItem(queueItem2)
	assert.NoError(t, err)

	prioritize := true
	queueItems, err := repo.GetQueueItemsByUserID(1, &prioritize)
	assert.NoError(t, err)
	assert.Len(t, queueItems, 2)
	assert.Equal(t, queueItem2.TrackURI, queueItems[0].TrackURI) // queueItem2 should come first due to higher weight
	assert.Equal(t, queueItem1.TrackURI, queueItems[1].TrackURI)
}

func TestGetQueueItemsBySessionIDByUserID(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := queue.NewGormQueueRepository(db)

	session := &models.Session{Slug: "slug1", HostID: 1}
	err = db.Create(session).Error
	assert.NoError(t, err)

	queueItem1 := &models.Queue{
		TrackURI:  "spotify:track:123",
		SessionID: session.ID,
		UserID:    1,
		Weight:    10,
	}
	queueItem2 := &models.Queue{
		TrackURI:  "spotify:track:456",
		SessionID: session.ID,
		UserID:    1,
		Weight:    20,
	}
	queueItem3 := &models.Queue{
		TrackURI:  "spotify:track:789",
		SessionID: session.ID,
		UserID:    2,
		Weight:    30,
	}

	err = repo.CreateQueueItem(queueItem1)
	assert.NoError(t, err)
	err = repo.CreateQueueItem(queueItem2)
	assert.NoError(t, err)
	err = repo.CreateQueueItem(queueItem3)
	assert.NoError(t, err)

	prioritize := true
	queueItems, err := repo.GetQueueItemsBySessionIDByUserID(session.ID, 1, &prioritize)
	assert.NoError(t, err)
	assert.Len(t, queueItems, 2)
	assert.Equal(t, queueItem2.TrackURI, queueItems[0].TrackURI) // queueItem2 should come first due to higher weight
	assert.Equal(t, queueItem1.TrackURI, queueItems[1].TrackURI)
}

func TestGetQueueItem(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := queue.NewGormQueueRepository(db)

	queueItem := &models.Queue{
		TrackURI:  "spotify:track:123",
		SessionID: 1,
		UserID:    1,
		Weight:    10,
	}

	err = repo.CreateQueueItem(queueItem)
	assert.NoError(t, err)

	retrievedQueueItem, err := repo.GetQueueItem(queueItem.ID)
	assert.NoError(t, err)
	assert.Equal(t, queueItem.TrackURI, retrievedQueueItem.TrackURI)
	assert.Equal(t, queueItem.SessionID, retrievedQueueItem.SessionID)
	assert.Equal(t, queueItem.UserID, retrievedQueueItem.UserID)
	assert.Equal(t, queueItem.Weight, retrievedQueueItem.Weight)
}

func TestUpdateQueueItem(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := queue.NewGormQueueRepository(db)

	queueItem := &models.Queue{
		TrackURI:  "spotify:track:123",
		SessionID: 1,
		UserID:    1,
		Weight:    10,
	}

	err = repo.CreateQueueItem(queueItem)
	assert.NoError(t, err)

	queueItem.Weight = 20
	err = repo.UpdateQueueItem(queueItem)
	assert.NoError(t, err)

	var updatedQueueItem models.Queue
	err = db.First(&updatedQueueItem, queueItem.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, queueItem.Weight, updatedQueueItem.Weight)
}

func TestDeleteQueueItem(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := queue.NewGormQueueRepository(db)

	queueItem := &models.Queue{
		TrackURI:  "spotify:track:123",
		SessionID: 1,
		UserID:    1,
		Weight:    10,
	}

	err = repo.CreateQueueItem(queueItem)
	assert.NoError(t, err)

	err = repo.DeleteQueueItem(queueItem.ID)
	assert.NoError(t, err)

	var deletedQueueItem models.Queue
	err = db.First(&deletedQueueItem, queueItem.ID).Error
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
