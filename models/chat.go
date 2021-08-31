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
	Participants       []ChatParticipant
	Messages           []Message
	DataSourceID       uint
	Tags               []Tag `gorm:"many2many:tag_chat;"`
}

func (Chat) TableName() string {
	return "chat"
}
