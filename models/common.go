package models

type Common struct {
	ID           uint   `gorm:"primaryKey"`
	Checked      bool   `gorm:"not null"`
	DeletedState string `gorm:"not null"`
}
