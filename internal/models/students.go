package models

import "time"

type Student struct {
	ID               uint      `gorm:"primaryKey"`
	Name             string    `gorm:"size:255;not null"`
	Email            string    `gorm:"uniqueIndex;size:255;not null"`
	IsPrefect        bool      `gorm:"not null"`
	IsLibrarian      bool      `gorm:"not null"`
	CurrentClassId   uint      `json:"current_class_id"`
	DateJoinedSchool time.Time `gorm:"not null"`
	Password         string    `gorm:"size:255;not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	User   User   `gorm:"foreignKey:UserID"`
	Parent Parent `gorm:"foreignKey:ParentID"`
}
