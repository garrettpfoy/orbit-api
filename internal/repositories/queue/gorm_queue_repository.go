package queue

import (
	"garrettpfoy/orbit-api/internal/models"

	"gorm.io/gorm"
)

type GormQueueRepository struct {
	db *gorm.DB
}

func NewGormQueueRepository(db *gorm.DB) *GormQueueRepository {
	return &GormQueueRepository{db: db}
}

func (r *GormQueueRepository) CreateQueueItem(queueItem *models.Queue) error {
	return r.db.Create(queueItem).Error
}

func (r *GormQueueRepository) GetQueueItemsBySessionID(sessionID uint, prioritize *bool) ([]models.Queue, error) {
	var queueItems []models.Queue
	query := r.db.Where("session_id = ?", sessionID).Preload("Session").Preload("User")
	if prioritize != nil && *prioritize {
		query = query.Order("weight DESC")
	}
	err := query.Find(&queueItems).Error
	return queueItems, err
}

func (r *GormQueueRepository) GetQueueItemsByUserID(userID uint, prioritize *bool) ([]models.Queue, error) {
	var queueItems []models.Queue
	query := r.db.Where("user_id = ?", userID).Preload("Session").Preload("User")
	if prioritize != nil && *prioritize {
		query = query.Order("weight DESC")
	}
	err := query.Find(&queueItems).Error
	return queueItems, err
}

func (r *GormQueueRepository) GetQueueItemsBySessionIDByUserID(sessionID, userID uint, prioritize *bool) ([]models.Queue, error) {
	var queueItems []models.Queue
	query := r.db.Where("session_id = ? AND user_id = ?", sessionID, userID).Preload("Session").Preload("User")
	if prioritize != nil && *prioritize {
		query = query.Order("weight DESC")
	}
	err := query.Find(&queueItems).Error
	return queueItems, err
}

func (r *GormQueueRepository) GetQueueItem(id uint) (*models.Queue, error) {
	var queueItem models.Queue
	err := r.db.Preload("Session").Preload("User").First(&queueItem, id).Error
	return &queueItem, err
}

func (r *GormQueueRepository) UpdateQueueItem(queueItem *models.Queue) error {
	return r.db.Save(queueItem).Error
}

func (r *GormQueueRepository) DeleteQueueItem(id uint) error {
	return r.db.Delete(&models.Queue{}, id).Error
}
