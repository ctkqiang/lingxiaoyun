package models

import "time"

type Students struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Email       string `gorm:"uniqueIndex;size:255;not null"`
	Class       string `gorm:"size:255;not null"`
	IsPrefect   bool   `gorm:"not null"`
	IsLibrarian bool   `gorm:"not null"`
	Password    string `gorm:"size:255;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
