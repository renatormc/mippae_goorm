package models

type Device struct {
	Common
	Folder      string
	Group       string
	DataSources []DataSource
}

func (Device) TableName() string {
	return "device"
}
