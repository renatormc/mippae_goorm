package models

type Contact struct {
	Common
	Name         string `gorm:"not null"`
	Source       string `gorm:"not null"`
	Entries      []ContactEntry
	DataSourceID uint
	Tags         []Tag `gorm:"many2many:tag_contact;"`
}

func (Contact) TableName() string {
	return "contact"
}
