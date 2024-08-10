package models

type Status struct {
	BaseModel
	Name string `gorm:"not null"`
}
