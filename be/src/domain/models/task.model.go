package models

import (
	"time"
)

type Task struct {
	BaseModel
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	ProjectID   uint   `gorm:"not null"`
	Status      string `gorm:"not null"`
	Deadline    *time.Time
}
