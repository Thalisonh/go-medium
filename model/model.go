package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	CreatedAt *time.Time `gorm:"<-:create"`
	Name      string     `gorm:"name" json:"name"`
	Author    string     `gorm:"author" json:"author"`
}
