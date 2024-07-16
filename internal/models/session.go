package models

import (
	"gorm.io/gorm"
)

// Session represents the sessions table
type Session struct {
	gorm.Model
	Slug   uint `gorm:"unique;not null"`
	HostID uint `gorm:"not null"`
	Host   User `gorm:"foreignKey:HostID"`
}
