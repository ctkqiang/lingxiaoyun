package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StudentID uint      `json:"student_id"`
	TeacherID uint      `json:"teacher_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`

	Student Student `gorm:"foreignKey:StudentID"`
	Teacher Teacher `gorm:"foreignKey:TeacherID"`
}
