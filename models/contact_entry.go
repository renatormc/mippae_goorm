package models

type ContactEntry struct {
	Common
	Category     string `gorm:"not null"`
	Value        string `gorm:"not null"`
	ContactID    uint
	DataSourceID uint
}

func (ContactEntry) TableName() string {
	return "contact_entry"
}
