package models

type ReadSource struct {
	common
	Folder        string
	Process       bool
	DataFile      string
	SourceType    string
	RegexSpiTools string
	ChatSource    string
}

func (ReadSource) TableName() string {
	return "read_source"
}
