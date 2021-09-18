package models

import (
	"time"
)

type Message struct {
	Common
	Timestamp              *time.Time
	Body                   string
	DeletedState           string
	ChatID                 uint
	FromIdentifier         string
	FromFriendlyIdentifier string
	FromName               string
	Color                  string
	PageRenderized         int
	Attachments            []File `gorm:"foreignKey:MessageID"`
	DataSourceID           uint
	AnaliseAttachmentTypes string `gorm:"default:''"`
}

func (Message) TableName() string {
	return "message"
}

func (m *Message) SetTimestampAsUnix(unixTime *int64) {
	if unixTime != nil {
		aux := time.Unix(*unixTime, 0)
		m.Timestamp = &aux
	}
}
