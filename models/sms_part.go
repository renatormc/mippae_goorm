package models

type SmsPart struct {
	Common
	Role       string
	Identifier string
	Name       string
	SmsID      uint
}

func (SmsPart) TableName() string {
	return "sms_part"
}
