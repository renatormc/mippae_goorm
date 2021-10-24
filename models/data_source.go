package models

type DataSource struct {
	Common
	SourceType    string
	Identifier    string `gorm:"not null"`
	Folder        string `gorm:"not null"`
	Process       bool   `gorm:"not null"`
	DataFile      string
	RegexSpiTools string
	ChatSource    string
	DeviceID      uint
	Device        Device `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (DataSource) TableName() string {
	return "data_source"
}
