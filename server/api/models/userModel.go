package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string  `json:"id" sql:"AUTO_INCREMENT" gorm:"primaryKey"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	UserName  string  `json:"userName" gorm:"unique"`
	Password  string  `json:"password"`
	Account   Account `json:"accounts" gorm:"foreignKey:UserId"`
}
