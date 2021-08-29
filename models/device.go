package models

type Device struct {
	common
	Folder string
	Group  string
}

func (Device) TableName() string {
	return "device"
}
