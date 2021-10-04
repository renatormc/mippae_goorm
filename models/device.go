package models

type Device struct {
	Common
	Identifier string `json:"identifier" gorm:"not null"`
	Folder     string `json:"-" gorm:"not null"`
	// Group       string       `json:"-"`
	DataSources []DataSource `json:"-"`
}

func (Device) TableName() string {
	return "device"
}
