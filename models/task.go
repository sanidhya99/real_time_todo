package models 

import(
	"gorm.io/gorm"
)


type TodoList struct {
	gorm.Model
	Name   string
	UserID uint
	Tasks  []Task
}

type Task struct{
	gorm.Model
	Description string `json:"description"`
	Completed bool `json:"completed"`
	TodoListID uint `json:"list"`
}