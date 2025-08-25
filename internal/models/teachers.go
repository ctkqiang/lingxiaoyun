package models

type Teacher struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"uniqueIndex"`
	User   User   `gorm:"constraint:OnDelete:CASCADE;"`
	Dept   string `gorm:"size:255"`
}
