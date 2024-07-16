package queue

import (
	"garrettpfoy/orbit-api/internal/models"
)

type QueueRepository interface {
	CreateQueueItem(queueItem *models.Queue) error
	GetQueueItemsBySessionID(sessionID uint) ([]models.Queue, error)
	GetQueueItem(id uint) (*models.Queue, error)
	UpdateQueueItem(queueItem *models.Queue) error
	DeleteQueueItem(id uint) error
}
