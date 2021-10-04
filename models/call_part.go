package models

type CallPart struct {
	ID           int
	Role         string `gorm:"not null"`
	Identifier   string `gorm:"not null"`
	Name         string `gorm:"not null"`
	CallID       uint
	DataSourceID uint
}

func (CallPart) TableName() string {
	return "call_part"
}
