package models

type Parent struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`

	User User `gorm:"foreignKey:UserID"`
}
