package models

type Teacher struct {
	ID     uint
	UserID uint
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Dept   string
}
