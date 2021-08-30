package models

import "time"

type Sms struct {
	Common
	Body      string
	Timestamp *time.Time
	Status    string
	Folder    string
	Parts     []SmsPart
	Tags      []Tag `gorm:"many2many:tag_sms;"`
}

func (Sms) TableName() string {
	return "sms"
}
