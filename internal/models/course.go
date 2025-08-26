package models

type Course struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TeacherID   uint   `json:"teacher_id"`

	Teacher Teacher `gorm:"foreignKey:TeacherID"`
}
