package business

import _ "gorm.io/gorm"

type Product struct {
	ID          int
	Title       string  `gorm:"size:255;not null"`
	Author      string  `gorm:"size:255;not null"`
	Category    string  `gorm:"size:255;not null"`
	Price       float64 `gorm:"not null"`
	Description string
	Language    string
	Cover       string
	Editorial   string
	Year        int
	Pages       int
}
