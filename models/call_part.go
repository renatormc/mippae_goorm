package models

type CallPart struct {
	Common
	Role       string `gorm:"not null"`
	Identifier string `gorm:"not null"`
	Name       string `gorm:"not null"`
	CallID     uint
	Call       Call
}

func (CallPart) TableName() string {
	return "call_part"
}
