package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type Task struct {
	gorm.Model
	Title       string
	Description string
	UserID      uint
}
