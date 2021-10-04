package models

type Tag struct {
	ID          uint
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Highlight   bool   `gorm:"not null,default:false"`
	Color       string `gorm:"not null,default:#000000"`
}

func (Tag) TableName() string {
	return "tag"
}
