package services

import (
	"fmt"
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

	userId, err := u.getUserIdByUsername(userName)

	if err != nil {
		return nil, fmt.Errorf("unable to fetch balance")
	}

	account := models.Account{}

	if err := u.db.Find(&account, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (u *AccountServices) TransferAmount(from string, to string, amount float64) error {
	toUserId, err := u.getUserIdByUsername(to)
	if err != nil {
		return fmt.Errorf("unable to find user")
	}

	fromUserId, err := u.getUserIdByUsername(from)

	if err != nil {
		return fmt.Errorf("unable to find user")
	}

	txnError := u.db.Transaction(func(tx *gorm.DB) error {
		toAcc := models.Account{}
		fromAcc := models.Account{}

		if err := tx.Find(&toAcc, "user_id = ?", toUserId).Error; err != nil {
			return fmt.Errorf("unable to find recipient's account")
		}

		if err := tx.Find(&fromAcc, "user_id = ?", fromUserId).Error; err != nil {
			return fmt.Errorf("unable to find sender's account")
		}

		toAcc.Balance += amount
		fromAcc.Balance -= amount

		if err := tx.Save(&toAcc).Error; err != nil {
			return fmt.Errorf("unable to update recipient's account")
		}

		if err := tx.Save(&fromAcc).Error; err != nil {
			return fmt.Errorf("unable to update sender's account")
		}

		return nil
	})

	if txnError != nil {
		return fmt.Errorf("unable to transfer amount")
	}

	return nil
}

func (u *AccountServices) getUserIdByUsername(username string) (string, error) {
	var user models.User
	if err := u.db.Find(&user, "user_name = ?", username).Error; err != nil {
		return "", err
	}

	if user.UserName == "" {
		return "", nil
	}

	return user.ID, nil
}
