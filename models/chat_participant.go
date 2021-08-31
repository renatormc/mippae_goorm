package models

type ChatParticipant struct {
	Common
	Identifier         string
	FriendlyIdentifier string
	Name               string
	Proprietary        bool
	Avatar             string
	ReadSourceID       uint
	ChatID             uint
}

func (ChatParticipant) TableName() string {
	return "participant"
}
