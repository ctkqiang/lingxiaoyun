package models

import "time"

type Attendance struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StudentID uint      `json:"student_id"`
	TeacherID uint      `json:"teacher_id"`
	Date      time.Time `json:"date"`
	Status    string    `gorm:"type:enum('present','absent','leave');not null" json:"status"`

	Student Student `gorm:"foreignKey:StudentID"`
	Teacher Teacher `gorm:"foreignKey:TeacherID"`
}
