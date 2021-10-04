package models

type Config struct {
	ID    uint
	Key   string `gorm:"not null"`
	Value string `gorm:"not null"`
}

func (Config) TableName() string {
	return "config"
}
