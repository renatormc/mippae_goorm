package models

import (
	"time"
)

type Chat struct {
	common
	Name               string
	Identifier         string
	FriendlyIdentifier string
	StartTime          time.Time
	LastActivity       time.Time
	NMessages          int
	Source             string
	Avatar             string
}

func (Chat) TableName() string {
	return "chat"
}
