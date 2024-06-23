package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskLog struct {
	gorm.Model
	TaskID      uint `gorm:"not null"`
	Description string
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	Duration    int       // Duration in minutes
}
