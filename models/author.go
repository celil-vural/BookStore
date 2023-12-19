// models/author.go
package models

import "gorm.io/gorm"

// Author model
type Author struct {
	gorm.Model
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
