package models

import (
	"time"
)

type Sms struct {
	Common
	Body         string `gorm:"not null"`
	Timestamp    *time.Time
	Status       string    `gorm:"not null"`
	Folder       string    `gorm:"not null"`
	Parts        []SmsPart `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataSourceID uint
	DataSource   DataSource
	Tags         []Tag `gorm:"many2many:tag_sms;"`
}

func (Sms) TableName() string {
	return "sms"
}
