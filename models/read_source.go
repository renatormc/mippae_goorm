package models

type DataSource struct {
	Common
	Identifier    string
	Folder        string
	Process       bool
	DataFile      string
	SourceType    string
	RegexSpiTools string
	ChatSource    string
	DeviceID      uint
}

func (DataSource) TableName() string {
	return "read_source"
}
