package database

import (
	"yukevent/config"
	"yukevent/etc"
	"yukevent/model"
)

func RegisterUser(user model.User) (*model.User, error) {

	hashedPassword, err := etc.HashingPassword(user.Password)
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

// func LoginUser(user mode.User) (*model.User, error) {

// }

func UpdateUser(id int, user model.User) (*model.User, error) {

	var users model.User
	err := config.DB.Find(&users, id).Updates(&user).Error
	if err != nil {
		return &model.User{}, err
	}
	return &user, err
}
