package models

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Tasks []Task `gorm:"foreignkey:StatusID"`
}
