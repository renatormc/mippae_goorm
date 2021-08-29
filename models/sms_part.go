package models

type SmsPart struct {
	common
	Role       string
	Identifier string
	Name       string
}

func (SmsPart) TableName() string {
	return "sms_part"
}
