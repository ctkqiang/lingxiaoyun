package models

import (
	"fmt"
	"time"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleTeacher Role = "teacher"
	RoleStudent Role = "student"
	RoleParent  Role = "parent"
)

func ParseRole(r string) (Role, error) {
	switch Role(r) {
	case RoleAdmin, RoleTeacher, RoleStudent, RoleParent:
		return Role(r), nil
	default:
		return "", fmt.Errorf("invalid role: %s", r)
	}
}

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"` // 永远不暴露
	Role      Role      `gorm:"type:varchar(20);not null" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
