package models

type Common struct {
	ID           uint `gorm:"primaryKey"`
	Checked      bool
	DeletedState string
}
