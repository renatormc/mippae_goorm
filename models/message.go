package models

import (
	"time"
)

type Message struct {
	common
	Timestamp    time.Time
	Body         string
	DeletedState string
}

func (Message) TableName() string {
	return "message"
}
