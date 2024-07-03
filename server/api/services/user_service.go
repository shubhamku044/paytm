package services

import (
	"fmt"
	"server/api/models"
	"server/api/utils"

	"gorm.io/gorm"
)

type UserServices struct {
	db *gorm.DB
}

func (u *UserServices) InitServices(db *gorm.DB) {
	u.db = db
	u.db.AutoMigrate(&models.User{})
}

func (u *UserServices) CreateUserService(user *models.User) (*models.User, error) {
	uuid := utils.GenerateUUID()
	user.ID = uuid
	hashedPwd, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}
	user.Password = hashedPwd

	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	accountBalance := utils.GenerateRandomNumber(1000, 10000)

	if err := u.db.Create(&models.Account{
		UserId:  user.ID,
		Balance: float64(accountBalance),
	}).Error; err != nil {
		fmt.Println("Unable to create Account")
		fmt.Println(err.Error())
	}

	return user, nil
}

func (u *UserServices) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := u.db.Find(&user, "user_name = ?", username).Error; err != nil {
		return nil, err
	}

	if user.UserName == "" {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (u *UserServices) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := u.db.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (u *UserServices) GetUserAccountDetails(userId string) (*models.Account, error) {
	var account models.Account
	if err := u.db.Find(&account, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	if account.UserId == "" {
		return nil, fmt.Errorf("account not found")
	}

	return &account, nil
}
