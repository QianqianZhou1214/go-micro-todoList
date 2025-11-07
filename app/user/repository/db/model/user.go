package model

import "gorm.io/gorm"

// define DB model
type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string
}
