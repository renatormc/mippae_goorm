package models

import (
	"database/sql"
	"time"
)

type Call struct {
	Common
	Type         sql.NullString
	Timestamp    *time.Time
	Duration     *time.Duration
	Parts        []CallPart
	DataSourceID uint
	DataSource   DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tags         []Tag      `gorm:"many2many:tag_call;"`
}

func (Call) TableName() string {
	return "call"
}
