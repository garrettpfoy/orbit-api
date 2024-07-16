package queue

import (
	"garrettpfoy/orbit-api/internal/models"

	"gorm.io/gorm"
)

type GormQueueRepository struct {
	DB *gorm.DB
}

func (r *GormQueueRepository) CreateQueueItem(queueItem *models.Queue) error {
	return r.DB.Create(queueItem).Error
}

func (r *GormQueueRepository) GetQueueItemsBySessionID(sessionID uint) ([]models.Queue, error) {
	var queueItems []models.Queue
	err := r.DB.Where("session_id = ?", sessionID).Find(&queueItems).Error
	return queueItems, err
}

func (r *GormQueueRepository) GetQueueItem(id uint) (*models.Queue, error) {
	var queueItem models.Queue
	err := r.DB.First(&queueItem, id).Error
	return &queueItem, err
}

func (r *GormQueueRepository) UpdateQueueItem(queueItem *models.Queue) error {
	return r.DB.Save(queueItem).Error
}

func (r *GormQueueRepository) DeleteQueueItem(id uint) error {
	return r.DB.Delete(&models.Queue{}, id).Error
}
