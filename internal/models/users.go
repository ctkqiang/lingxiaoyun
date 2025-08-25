package models

import "time"

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleTeacher Role = "teacher"
	RoleStudent Role = "student"
	RoleParent  Role = "parent"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"uniqueIndex;size:255;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      Role   `gorm:"type:enum('admin','teacher','student','parent');not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
