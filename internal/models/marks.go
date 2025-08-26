package models

type Mark struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	StudentID uint    `json:"student_id"`
	TeacherID uint    `json:"teacher_id"`
	Subject   string  `json:"subject"`
	Term      string  `json:"term"`
	Score     float64 `json:"score"`

	Student Student `gorm:"foreignKey:StudentID"`
	Teacher Teacher `gorm:"foreignKey:TeacherID"`
}
