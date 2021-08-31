package models

import "time"

type File struct {
	Common
	Size          int
	Filename      string
	OriginalPath  string
	ExtractedPath string
	ConvertedPath string
	Extension     string
	ContentType   string
	CreationTime  *time.Time
	ModifyTime    *time.Time
	AccessTime    *time.Time
	Sha256        string
	Md5           string
	Type          string
	Corrupted     bool
	MessageID     uint
	ReadSourceID  uint
	Tags          []Tag `gorm:"many2many:tag_file;"`
}

func (File) TableName() string {
	return "file"
}
