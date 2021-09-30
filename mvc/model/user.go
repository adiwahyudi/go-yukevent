package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           int    `json:"id" form:"id" gorm:"primary_key"`
	Username     string `json:"username" form:"username" gorm:"unique"`
	Email        string `json:"email" form:"email" gorm:"unique"`
	Password     string `json:"password" form:"password"`
	Name         string `json:"name" form:"name"`
	Dob          string `json:"dob" form:"dob"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
	Photo        string `json:"photo" form:"photo"`
}
