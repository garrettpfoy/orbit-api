package vote

import (
	"garrettpfoy/orbit-api/internal/models"
)

type VoteRepository interface {
	CreateVote(vote *models.Vote) error
	GetVotesByQueueID(queueID uint) ([]models.Vote, error)
	GetVote(id uint) (*models.Vote, error)
	UpdateVote(vote *models.Vote) error
	DeleteVote(id uint) error
}
