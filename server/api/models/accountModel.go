package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserId  string  `json:"user_id" gorm:"primaryKey"`
	Balance float64 `json:"balance" gorm:"default:10000"`
}
