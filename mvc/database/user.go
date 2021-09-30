package database

import (
	"yukevent/config"
	"yukevent/helper"
	"yukevent/model"
)

func RegisterUser(user model.User) (*model.User, error) {

	hashedPassword, err := helper.HashingPassword(user.Password)
	user.Password = hashedPassword
	if err != nil {
		return &model.User{}, err
	}

	err = config.DB.Save(&user).Error

	if err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func LoginUser(email, password string) (*model.User, error) {

	var users model.User

	err := config.DB.Where("email = ?", email).First(&users).Error

	if err != nil {
		return &model.User{}, err
	}

	err = helper.CheckPasswordHash(password, users.Password)

	if err != nil {
		return &model.User{}, err
	}

	return &users, nil
}

func GetUser() (*[]model.User, error) {

	var users []model.User

	err := config.DB.Find(&users).Error

	if err != nil {
		return &[]model.User{}, err
	}
	return &users, nil
}

// func UpdateUser(id int, user model.User) (*model.User, error) {

// 	var users model.User
// 	err := config.DB.Find(&users, id).Updates(&user).Error
// 	if err != nil {
// 		return &model.User{}, err
// 	}
// 	return &user, err
// }
