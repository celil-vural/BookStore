// models/author.go
package models

import (
	"gorm.io/gorm"
)

// Author model
type Author struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;type:varchar(50)"`
	Surname string `json:"Surname" gorm:"not null;type:varchar(50)"`
}
