package models

type ChatParticipant struct {
	ID                 uint `gorm:"primaryKey"`
	Identifier         string
	FriendlyIdentifier string
	Name               string
	Proprietary        bool
	Avatar             string
	DataSourceID       uint
	Chats              []Chat `gorm:"many2many:chat_participant;"`
}

func (ChatParticipant) TableName() string {
	return "participant"
}
