package models

type DataSource struct {
	Common
	Identifier    string `json:"identifier"`
	Folder        string `json:"-"`
	Process       bool   `json:"process"`
	DataFile      string `json:"ufed_xml_file"`
	SourceType    string `json:"type"`
	RegexSpiTools string `json:"spi_regex"`
	ChatSource    string `json:"sqlite_chat_name"`
	DeviceID      uint   `json:"-"`
}

func (DataSource) TableName() string {
	return "data_source"
}
