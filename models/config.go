package models

type Config struct {
	ID    uint
	Key   string
	Value string
}

func (Config) TableName() string {
	return "config"
}
