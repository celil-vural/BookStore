package models

import (
	"gorm.io/gorm"
	"time"
)

// Book model
type Book struct {
	gorm.Model
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	// Relationship
	Genres []Genre `gorm:"many2many:book_genres;"`
	Author *Author `gorm:"foreignkey:author_id"`
}
