package models

import "database/sql"

type ChatParticipant struct {
	Common
	Identifier         string `gorm:"not null"`
	FriendlyIdentifier string `gorm:"not null"`
	Name               string `gorm:"not null"`
	Proprietary        bool   `gorm:"not null"`
	Avatar             sql.NullString
	ChatID             uint
	Chat               Chat
}

func (ChatParticipant) TableName() string {
	return "chat_participant"
}
