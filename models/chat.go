package models

import (
	"database/sql"
	"time"
)

type Chat struct {
	Common
	Name               string `gorm:"not null"`
	Identifier         string `gorm:"not null"`
	FriendlyIdentifier string `gorm:"not null"`
	StartTime          *time.Time
	LastActivity       *time.Time
	NMessages          int    `gorm:"not null"`
	Source             string `gorm:"not null"`
	Avatar             sql.NullString
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
