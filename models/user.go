package models

type User struct {
	ID 			uint64	`gorm:"unique"`
	Name 		string	`json:"name"`
	Email 		string 	`json:"email" gorm:"unique"`
	Password 	[]byte	`json:"-"` // don't return the value
	Posts 		[]Post	`json:"posts" gorm:"foreignKey:UserId"`
}
