package models

type Cart struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"uniqueIndex"`
	Items  []Item `gorm:"many2many:cart_items"`
}