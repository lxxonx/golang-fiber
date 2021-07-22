package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID 			uint64	`gorm:"primaryKey"`
	Name 		string	`json:"name"`
	Email 		string 	`json:"email" gorm:"unique"`
	Password 	[]byte	`json:"-"` // don't return the value
	Posts 		[]Post	`json:"posts" gorm:"foreignKey:UserId"`
}
