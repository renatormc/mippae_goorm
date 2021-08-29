package models

type ChatParticipant struct {
	common
	Identifier         string
	FriendlyIdentifier string
	Name               string
	Proprietary        bool
	Avatar             string
}

func (ChatParticipant) TableName() string {
	return "participant"
}
