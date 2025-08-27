package models

import "time"

type Student struct {
	ID               uint
	UserID           uint
	User             User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CurrentClassId   uint
	DateJoinedSchool time.Time
	Password         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
