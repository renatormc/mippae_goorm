package models

type UserAccount struct {
	Common
	Name         string
	Username     string
	ServiceType  string
	Password     string
	DataSourceID uint
}

func (UserAccount) TableName() string {
	return "user_account"
}
