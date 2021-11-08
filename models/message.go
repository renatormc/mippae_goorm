package models

import (
	"time"
)

type Message struct {
	Common
	Timestamp              *time.Time
	Body                   string `gorm:"not null"`
	DeletedState           string `gorm:"not null"`
	ChatID                 uint
	Color                  string `gorm:"not null,default:'#000000'"`
	PageRenderized         int64  `gorm:"not null,default:-1"`
	Attachments            []File `gorm:"foreignKey:MessageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AnaliseAttachmentTypes string `gorm:"not null,default:''"`
	FromID                 uint
	From                   ChatParticipant `gorm:"foreignKey:FromID;"`
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
