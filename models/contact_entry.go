package models

type ContactEntry struct {
	Common
	Category     string `gorm:"not null"`
	Value        string `gorm:"not null"`
	ContactID    uint
	Contact      Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataSourceID uint
	DataSource   DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ContactEntry) TableName() string {
	return "contact_entry"
}
