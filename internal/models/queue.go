package models

import "gorm.io/gorm"

// Queue represents the queue table
type Queue struct {
	gorm.Model
	TrackURI  string  `gorm:"not null"`
	SessionID uint    `gorm:"not null"`
	Session   Session `gorm:"foreignKey:SessionID"`
	Weight    int     `gorm:"default:0;not null"`
}
