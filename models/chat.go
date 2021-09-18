package models

import (
	"time"
)

type Chat struct {
	Common
	Name               string
	Identifier         string
	FriendlyIdentifier string
	StartTime          *time.Time
	LastActivity       *time.Time
	NMessages          int
	Source             string
	Avatar             string
	Participants       []ChatParticipant `gorm:"foreignKey:ChatID"`
	Messages           []Message
	DataSourceID       uint
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
