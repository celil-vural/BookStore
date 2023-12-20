package models

import "gorm.io/gorm"

// Book model
type Book struct {
	gorm.Model
	Title    string  `json:"title" gorm:"unique"`
	Price    float64 `json:"price" gorm:"not null"`
	AuthorID int     `json:"author_id" gorm:"not null"`
	// Relationship
	Genres []*Genre `json:"genres" gorm:"many2many:book_genres;"`
	Author *Author  `gorm:"foreignKey:AuthorID"`
}
