package models

import (
	"time"
)

type Chat struct {
	Common
	Name               string `gorm:"not null"`
	Identifier         string `gorm:"not null"`
	FriendlyIdentifier string `gorm:"not null"`
	StartTime          *time.Time
	LastActivity       *time.Time
	NMessages          int64  `gorm:"not null, default: 0"`
	Source             string `gorm:"not null"`
	Avatar             string
	Participants       []ChatParticipant `gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Messages           []Message         `gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataSourceID       uint
	DataSource         DataSource
	Tags               []Tag `gorm:"many2many:tag_cha;"`
}

func (Chat) TableName() string {
	return "chat"
}

func (c *Chat) SetStartTimeAsUnix(unixTime *int64) {
	if unixTime != nil {
		creation := time.Unix(*unixTime, 0)
		c.StartTime = &creation
	}
}
