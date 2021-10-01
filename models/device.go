package models

type Device struct {
	Common
	Identifier  string       `json:"identifier"`
	Folder      string       `json:"-"`
	Group       string       `json:"-"`
	DataSources []DataSource `json:"-"`
}

func (Device) TableName() string {
	return "device"
}
