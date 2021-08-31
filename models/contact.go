package models

type Contact struct {
	Common
	Name         string
	Source       string
	Entries      []ContactEntry
	DataSourceID uint
	Tags         []Tag `gorm:"many2many:tag_contact;"`
}

func (Contact) TableName() string {
	return "contact"
}
