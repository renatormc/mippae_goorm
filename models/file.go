package models

import (
	"database/sql"
	"time"
)

type File struct {
	Common
	Size          sql.NullInt64
	Filename      string `gorm:"not null"`
	OriginalPath  sql.NullString
	ExtractedPath string `gorm:"not null"`
	ConvertedPath sql.NullString
	Extension     string `gorm:"not null, default: 'file'"`
	ContentType   sql.NullString
	CreationTime  *time.Time
	ModifyTime    *time.Time
	AccessTime    *time.Time
	Sha256        sql.NullString
	Md5           sql.NullString
	Type          sql.NullString
	Corrupted     bool `gorm:"not null"`
	MessageID     sql.NullInt64
	DataSourceID  uint
	Tags          []Tag `gorm:"many2many:tag_file;"`
}

func (File) TableName() string {
	return "file"
}
