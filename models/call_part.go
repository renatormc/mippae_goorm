package models

type CallPart struct {
	ID         int
	Role       string
	Identifier string
	Name       string
}

func (CallPart) TableName() string {
	return "call_part"
}
