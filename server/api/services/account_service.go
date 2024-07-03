package services

import (
	"server/api/models"

	"gorm.io/gorm"
)

type AccountServices struct {
	db *gorm.DB
}

func (u *AccountServices) InitServices(db *gorm.DB) {
	u.db = db
	u.db.AutoMigrate(&models.Account{})
}

func (u *AccountServices) GetBalanceByUserName(userName string) (*models.Account, error) {
	user := models.User{}
	if err := u.db.Find(&user, "user_name = ?", userName).Error; err != nil {
		return nil, err
	}

	if user.UserName == "" {
		return nil, nil
	}

	account := models.Account{}

	if err := u.db.Find(&account, "user_id = ?", user.ID).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
