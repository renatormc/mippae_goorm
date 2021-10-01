package models

type ChatParticipant struct {
	ID                 uint `gorm:"primaryKey"`
	Identifier         string
	FriendlyIdentifier string
	Name               string
	Proprietary        bool
	Avatar             string
	DataSourceID       uint
	ChatID             uint
}

func (ChatParticipant) TableName() string {
	return "chat_participant"
}
