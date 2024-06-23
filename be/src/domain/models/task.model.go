package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      uint `gorm:"not null"`
	ProjectID   uint
	Deadline    *time.Time
	StatusID    uint      `gorm:"not null"`
	TaskLogs    []TaskLog `gorm:"foreignkey:TaskID"`
}
