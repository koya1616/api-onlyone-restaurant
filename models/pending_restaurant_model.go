package models

type PendingRestaurant struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Information string `gorm:"not null"`
	Isapproved  bool   `gorm:"not null;default:false"`
}
