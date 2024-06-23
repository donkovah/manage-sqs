package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
}
