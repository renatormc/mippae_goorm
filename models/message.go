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
	Color                  string `gorm:"not null,default:'#000000'"`
	PageRenderized         sql.NullInt64
	Attachments            []File `gorm:"foreignKey:MessageID"`
	DataSourceID           uint
	DataSource             DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AnaliseAttachmentTypes string     `gorm:"not null,default:''"`
	FromID                 uint
	From                   ChatParticipant `gorm:"foreignKey:FromID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
