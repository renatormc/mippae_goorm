package models

type CallPart struct {
	ID           int
	Role         string `gorm:"not null"`
	Identifier   string `gorm:"not null"`
	Name         string `gorm:"not null"`
	CallID       uint
	Call         Call `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataSourceID uint
	DataSource   DataSource `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (CallPart) TableName() string {
	return "call_part"
}
