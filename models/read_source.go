package models

type ReadSource struct {
	Common
	Folder        string
	Process       bool
	DataFile      string
	SourceType    string
	RegexSpiTools string
	ChatSource    string
	DeviceID      uint
}

func (ReadSource) TableName() string {
	return "read_source"
}
