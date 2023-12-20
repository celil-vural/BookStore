package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User model
type User struct {
	ID         string    `json:"id" gorm:"type:varchar(50);primaryKey"`
	Username   string    `json:"username" gorm:"unique"`
	Email      string    `json:"email" gorm:"unique"`
	Password   string    `json:"password" gorm:"not null"`
	Created_At time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	Updated_At time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
	Deleted_At time.Time `json:"deleted_at" gorm:"default:null"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
