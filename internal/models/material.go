package models

type Material struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	CourseID uint   `json:"course_id"`
	Title    string `json:"title"`
	FileURL  string `json:"file_url"`

	Course Course `gorm:"foreignKey:CourseID"`
}
