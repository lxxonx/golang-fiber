package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID 			uint64 	`gorm:"primaryKey"`
	Title		string	`json:"title"`
	Text		string	`json:"text"`
	UserId		uint64	`json:"userId"`
	Music 		[]uint8	`json:"music"`
}