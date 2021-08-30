package models

type ContactEntry struct {
	Common
	Category  string
	Value     string
	ContactID uint
}

func (ContactEntry) TableName() string {
	return "contact_entry"
}
