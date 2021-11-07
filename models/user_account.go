package models

type UserAccount struct {
	Common
	Name         string `gorm:"not null"`
	Username     string `gorm:"not null"`
	ServiceType  string `gorm:"not null"`
	Password     string `gorm:"not null"`
	DataSourceID uint
	DataSource   DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserAccount) TableName() string {
	return "user_account"
}
