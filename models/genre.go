// models/genre.go
package models

import "gorm.io/gorm"

// Genre model
type Genre struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// BookGenre relation model
type BookGenre struct {
	gorm.Model
	BookID  uint `gorm:"primaryKey"`
	GenreID uint `gorm:"primaryKey"`
}
