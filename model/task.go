package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID      int
	Content string
}
