package models

type Tag struct {
	ID          uint
	Name        string
	Description string
	Highlight   bool
	Color       string
}

func (Tag) TableName() string {
	return "tag"
}
