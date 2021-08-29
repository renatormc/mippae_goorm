package models

type ContactEntry struct {
	common
	Category string
	Value    string
}

func (ContactEntry) TableName() string {
	return "contact_entry"
}
