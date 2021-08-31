package models

import "time"

type Call struct {
	Common
	Type         string
	Timestamp    *time.Time
	Duration     *time.Duration
	Parts        []CallPart
	ReadSourceID uint
	Tags         []Tag `gorm:"many2many:tag_call;"`
}

func (Call) TableName() string {
	return "call"
}
