package entity

import "gorm.io/gorm"

// Todo represent model of todo
type Todo struct {
	gorm.Model
	TaskName    string
	Description string
	IsDone      bool
}
