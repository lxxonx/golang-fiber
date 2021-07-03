package models

type Post struct {
	ID 			uint 	`json:"id" gorm:"unique"`
	Title		string	`json:"title"`
	Text		string	`json:"text"`
}