package models

import (
	"database/sql"
	"time"
)

type File struct {
	Common
	Size          int64
	Filename      string `gorm:"not null"`
	OriginalPath  string
	ExtractedPath string `gorm:"not null"`
	ConvertedPath string
	Extension     string `gorm:"not null, default: 'file'"`
	ContentType   string
	CreationTime  *time.Time
	ModifyTime    *time.Time
	AccessTime    *time.Time
	Sha256        string
	Md5           string
	Type          string
	Corrupted     bool `gorm:"not null"`
	MessageID     sql.NullInt64
	Message       Message
	DataSourceID  uint
	DataSource    DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tags          []Tag      `gorm:"many2many:tag_file;"`
}

func (File) TableName() string {
	return "file"
}
