package models

import (
	"time"
)

type Message struct {
	Common
	Timestamp    *time.Time
	Body         string
	DeletedState string
	ChatID       uint
}

func (Message) TableName() string {
	return "message"
}
