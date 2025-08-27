package models

type Parent struct {
	ID     uint
	UserID uint
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name   string
	Email  string
	Phone  string
}
