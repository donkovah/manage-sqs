package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	ProjectID   uint   `gorm:"not null"`
	Status      string `gorm:"not null"`
	Deadline    *time.Time
	TaskLogs    []TaskLog `gorm:"foreignkey:TaskID"`
}
