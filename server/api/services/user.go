package services

import (
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
