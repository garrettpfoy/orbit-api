package models

import "gorm.io/gorm"

// Vote represents the votes table
type Vote struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"foreignKey:UserID"`
	QueueID  uint   `gorm:"not null"`
	Queue    Queue  `gorm:"foreignKey:QueueID"`
	VoteType string `gorm:"not null"`
}
