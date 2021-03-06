package models

type SmsPart struct {
	Common
	Role       string `gorm:"not null"`
	Identifier string `gorm:"not null"`
	Name       string `gorm:"not null"`
	SmsID      uint
	Sms        Sms
}

func (SmsPart) TableName() string {
	return "sms_part"
}
