package models

type Project struct {
	BaseModel
	Name        string `json:"name" validate:"required,min=2,max=100" gorm:"not null"`
	Description string `json:"description" validate:"max=500"`
	Tasks       []Task `gorm:"foreignkey:ProjectID"`
}
