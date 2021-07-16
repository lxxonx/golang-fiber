package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID 			uint64 	`gorm:"unique"`
	Title		string	`json:"title"`
	Text		string	`json:"text"`
	UserId		uint64	`json:"userId"`
}