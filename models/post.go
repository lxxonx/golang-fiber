package models

type Post struct {
	ID 			uint64 	`json:"id" gorm:"unique"`
	Title		string	`json:"title"`
	Text		string	`json:"text"`
	UserId		uint64	`json:"userId"`
}