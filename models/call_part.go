package models

type CallPart struct {
	ID         int
	Role       string
	Identifier string
	Name       string
	CallID     uint
}

func (CallPart) TableName() string {
	return "call_part"
}
