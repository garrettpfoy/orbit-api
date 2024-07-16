package vote

import (
	"garrettpfoy/orbit-api/internal/models"

	"gorm.io/gorm"
)

type GormVoteRepository struct {
	DB *gorm.DB
}

func (r *GormVoteRepository) CreateVote(vote *models.Vote) error {
	return r.DB.Create(vote).Error
}

func (r *GormVoteRepository) GetVotesByQueueID(queueID uint) ([]models.Vote, error) {
	var votes []models.Vote
	err := r.DB.Where("queue_id = ?", queueID).Find(&votes).Error
	return votes, err
}

func (r *GormVoteRepository) GetVote(id uint) (*models.Vote, error) {
	var vote models.Vote
	err := r.DB.First(&vote, id).Error
	return &vote, err
}

func (r *GormVoteRepository) UpdateVote(vote *models.Vote) error {
	return r.DB.Save(vote).Error
}

func (r *GormVoteRepository) DeleteVote(id uint) error {
	return r.DB.Delete(&models.Vote{}, id).Error
}
