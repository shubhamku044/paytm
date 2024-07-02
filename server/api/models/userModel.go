package models

type User struct {
	ID        string `json:"id" sql:"AUTO_INCREMENT" gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName" gorm:"unique"`
	Password  string `json:"password"`
}
