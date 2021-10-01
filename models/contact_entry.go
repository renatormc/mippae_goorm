package models

type ContactEntry struct {
	Common
	Category     string
	Value        string
	ContactID    uint
	DataSourceID uint
}

func (ContactEntry) TableName() string {
	return "contact_entry"
}
