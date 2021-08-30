package models

type Device struct {
	Common
	Folder      string
	Group       string
	ReadSources []ReadSource
}

func (Device) TableName() string {
	return "device"
}
