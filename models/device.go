package models

type Device struct {
	Common
	Identifier  string
	Folder      string
	Group       string
	DataSources []DataSource
}

func (Device) TableName() string {
	return "device"
}
