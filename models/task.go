package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`

	UserId uint `json:"user_id"`

	// User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
