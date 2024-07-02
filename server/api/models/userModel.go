package models

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName" gorm:"unique"`
	Password  string `json:"password"`
}
