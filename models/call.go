package models

import (
	"time"
)

type Call struct {
	Common
	Type         string
	Timestamp    *time.Time
	Duration     *time.Duration
	Parts        []CallPart `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataSourceID uint
	DataSource   DataSource
	Tags         []Tag `gorm:"many2many:tag_call;"`
}

func (Call) TableName() string {
	return "call"
}
