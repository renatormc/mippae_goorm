package models

type CallPart struct {
	ID           int
	Role         string
	Identifier   string
	Name         string
	CallID       uint
	DataSourceID uint
}

func (CallPart) TableName() string {
	return "call_part"
}
