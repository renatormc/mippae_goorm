package models

type ContactEntry struct {
	Common
	Category  string `gorm:"not null"`
	Value     string `gorm:"not null"`
	ContactID uint
	Contact   Contact
}

func (ContactEntry) TableName() string {
	return "contact_entry"
}
