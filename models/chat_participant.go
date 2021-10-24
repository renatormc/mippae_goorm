package models

import "database/sql"

type ChatParticipant struct {
	ID                 uint   `gorm:"primaryKey"`
	Identifier         string `gorm:"not null"`
	FriendlyIdentifier string `gorm:"not null"`
	Name               string `gorm:"not null"`
	Proprietary        bool   `gorm:"not null"`
	Avatar             sql.NullString
	DataSourceID       uint
	DataSource         DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ChatID             uint
}

func (ChatParticipant) TableName() string {
	return "chat_participant"
}
