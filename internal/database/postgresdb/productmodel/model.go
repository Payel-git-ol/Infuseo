package productmodel

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
}
