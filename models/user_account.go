package models

import "database/sql"

type UserAccount struct {
	Common
	Name         string `gorm:"not null"`
	Username     string `gorm:"not null"`
	ServiceType  string `gorm:"not null"`
	Password     sql.NullString
	DataSourceID uint
}

func (UserAccount) TableName() string {
	return "user_account"
}
