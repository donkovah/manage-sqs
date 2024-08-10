package models

type Comment struct {
	BaseModel
	Name       string `gorm:"not null"`
	TimelineID int
	Timeline   Timeline
}
