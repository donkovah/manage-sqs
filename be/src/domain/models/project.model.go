package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	UserID      uint   `gorm:"not null"`
	Tasks       []Task `gorm:"foreignkey:ProjectID"`
}
