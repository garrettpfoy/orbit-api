package queue

import (
	"garrettpfoy/orbit-api/internal/models"
)

type QueueRepository interface {
	// CreateQueueItem validates and creates a new queue item in the database
	CreateQueueItem(queueItem *models.Queue) error
	// GetQueueItemsBySessionID retrieves all queue items in a session by the session ID
	// If prioritize is true, the queue items are sorted by weight in descending order
	GetQueueItemsBySessionID(sessionID uint, prioritize *bool) ([]models.Queue, error)
	// GetQueueItemsByUserID retrieves all queue items in a session by the user ID
	// If prioritize is true, the queue items are sorted by weight in descending order
	GetQueueItemsByUserID(userID uint, prioritize *bool) ([]models.Queue, error)
	// GetQueueItemsBySessionIDByUserID retrieves all queue items in a session by the session ID and user ID
	// If prioritize is true, the queue items are sorted by weight in descending order
	GetQueueItemsBySessionIDByUserID(sessionID, userID uint, prioritize *bool) ([]models.Queue, error)
	// GetQueueItem retrieves a queue item from the database by its ID
	GetQueueItem(id uint) (*models.Queue, error)
	// UpdateQueueItem validates and updates a queue item in the database
	UpdateQueueItem(queueItem *models.Queue) error
	// DeleteQueueItem deletes a queue item from the database by its ID
	DeleteQueueItem(id uint) error
}
