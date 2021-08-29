package models

type Contact struct {
	common
	Name   string
	Source string
}

func (Contact) TableName() string {
	return "contact"
}
