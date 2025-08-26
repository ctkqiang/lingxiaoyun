package models

type Class struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	TeacherID uint   `json:"teacher_id"`

	Teacher Teacher `gorm:"foreignKey:TeacherID"`
}
