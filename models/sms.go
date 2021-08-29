package models

import "time"

type Sms struct {
	common
	Body         string
	Timestamp    time.Time
	Status       string
	Folder       string
	DeletedState string
}

func (Sms) TableName() string {
	return "sms"
}
