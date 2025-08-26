package models

import "time"

type Leave struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	StudentID  uint      `json:"student_id"`
	ParentID   uint      `json:"parent_id"`
	Reason     string    `json:"reason"`
	Status     string    `gorm:"type:enum('pending','approved','rejected');default:'pending'" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	ApprovedBy *uint     `json:"approved_by"`

	Student  Student `gorm:"foreignKey:StudentID"`
	Parent   Parent  `gorm:"foreignKey:ParentID"`
	Approver *User   `gorm:"foreignKey:ApprovedBy"`
}
