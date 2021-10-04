package models

import (
	"database/sql"
	"time"
)

type Message struct {
	Common
	Timestamp              *time.Time
	Body                   string `gorm:"not null"`
	DeletedState           string `gorm:"not null"`
	ChatID                 uint
	FromIdentifier         string `gorm:"not null"`
	FromFriendlyIdentifier string `gorm:"not null"`
	FromName               string `gorm:"not null"`
	Color                  string `gorm:"not null,default:'#000000'"`
	PageRenderized         sql.NullInt64
	Attachments            []File `gorm:"foreignKey:MessageID"`
	DataSourceID           uint
	AnaliseAttachmentTypes string `gorm:"not null,default:''"`
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
