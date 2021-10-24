package models

type SmsPart struct {
	Common
	Role         string `gorm:"not null"`
	Identifier   string `gorm:"not null"`
	Name         string `gorm:"not null"`
	SmsID        uint
	Sms          Sms `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataSourceID uint
	DataSource   DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (SmsPart) TableName() string {
	return "sms_part"
}
