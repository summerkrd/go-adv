package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	SessionID string `gorm:"uniqueIndex;not null"`
	Phone     string `gorm:"not null"`
	Code      int    `gorm:"not null"`
	ExpiresAt time.Time
	Verified  bool `gorm:"default:false"`
}
