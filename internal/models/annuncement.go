package models

import "time"

type Announcement struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Audience  string    `gorm:"type:enum('all','teachers','parents','students');default:'all'" json:"audience"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:CreatedBy"`
}
