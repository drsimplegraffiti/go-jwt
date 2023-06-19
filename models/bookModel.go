package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string
	Author string
	Year string
	ISBN string
	UserID  uint // Foreign key referencing the User ID
	User User // The user that owns this book, coming from the User ID
}