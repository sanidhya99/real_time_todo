package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `gorm:"unique" json:"mail"`   // Correct format
	Password string    `json:"password"`
	TodoLists []TodoList
}
