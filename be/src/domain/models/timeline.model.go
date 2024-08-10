package models

type Timeline struct {
	BaseModel
	Name        string `gorm:"not null"`
	Description string
	Comment     []Comment
}
