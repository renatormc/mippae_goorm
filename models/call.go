package models

import "time"

type Call struct {
	Common
	Type      string
	Timestamp *time.Time
	Duration  *time.Duration
	Parts     []CallPart
}

func (Call) TableName() string {
	return "call"
}
