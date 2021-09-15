package models

import (
	"time"
)

type Message struct {
	Common
	Timestamp      *time.Time
	Body           string
	DeletedState   string
	ChatID         uint
	FromID         uint
	From           ChatParticipant
	Color          string
	PageRenderized int
	Attachments    []File
	DataSourceID   uint
}

func (Message) TableName() string {
	return "message"
}
