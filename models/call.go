package models

import "time"

type Call struct {
	common
	Type      string
	Timestamp time.Time
	Duration  time.Duration
}

func (Call) TableName() string {
	return "call"
}
