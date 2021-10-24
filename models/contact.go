package models

type Contact struct {
	Common
	Name         string `gorm:"not null"`
	Source       string `gorm:"not null"`
	Entries      []ContactEntry
	DataSourceID uint
	DataSource   DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tags         []Tag      `gorm:"many2many:tag_contact;"`
}

func (Contact) TableName() string {
	return "contact"
}
