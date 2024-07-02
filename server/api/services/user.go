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

func (u *UserServices) InitUserServices(db *gorm.DB) {
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
	return user, nil
}

func (u *UserServices) GetUserByUsername(username string) (*models.User, error) {
	fmt.Println(username)
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
