// models/genre.go
package models

import (
	"gorm.io/gorm"
	"time"
)

// Genre model
type Genre struct {
	GenreID    uint      `json:"genre_id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"unique"`
	Created_At time.Time `json:"created_at" gorm:"not null"`
	Updated_At time.Time `json:"updated_at" gorm:"not null"`
	Deleted_At time.Time `json:"deleted_at" gorm:"default:null"`
}

// BookGenre relation model
type BookGenre struct {
	gorm.Model
	BookID  uint `gorm:"primaryKey"`
	GenreID uint `gorm:"primaryKey"`
}
