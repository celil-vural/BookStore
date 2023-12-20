// models/genre.go
package models

import (
	"gorm.io/gorm"
)

// Genre model
type Genre struct {
	gorm.Model
	Name  string  `json:"name" gorm:"unique"`
	Books []*Book `gorm:"many2many:book_genres;"`
}
