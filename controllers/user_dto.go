package controllers

import (
	"time"
	"yukevent/model"
)

// Login Request and Response
type RequestUserRegister struct {
	Username     string `json:"username" form:"username"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Name         string `json:"name" form:"name"`
	Dob          string `json:"dob" form:"dob"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
	Photo        string `json:"photo" form:"photo"`
}

func (req *RequestUserRegister) toModel() model.User {
	return model.User{
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password,
		Name:         req.Name,
		Dob:          req.Dob,
		Phone_Number: req.Phone_Number,
		Photo:        req.Photo,
	}
}

type ResponseUser struct {
	ID           int       `json:"id" form:"id"`
	Username     string    `json:"username" form:"username"`
	Email        string    `json:"email" form:"email"`
	Name         string    `json:"name" form:"name"`
	Dob          string    `json:"dob" form:"dob"`
	Phone_Number string    `json:"phone_number" form:"phone_number"`
	Photo        string    `json:"photo" form:"photo"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at" form:"deleted_at"`
}

func newResponse(mdlUsers model.User) ResponseUser {
	return ResponseUser{
		ID:           mdlUsers.ID,
		Username:     mdlUsers.Username,
		Email:        mdlUsers.Email,
		Name:         mdlUsers.Name,
		Dob:          mdlUsers.Dob,
		Phone_Number: mdlUsers.Phone_Number,
		Photo:        mdlUsers.Photo,
		CreatedAt:    mdlUsers.CreatedAt,
		UpdatedAt:    mdlUsers.UpdatedAt,
		DeletedAt:    mdlUsers.DeletedAt.Time,
	}
}

// Login Request
type RequestUserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// type ResponeLoginUser struct {
// 	ID       int    `json:"id" form:"id"`
// 	Username string `json:"username" form:"username"`
// 	Email    string `json:"email" form:"email"`
// 	Name     string `json:"name" form:"name"`
// 	Token    string `json:"jwt" form:"jwt"`
// }

// func newResponseLoginUser(mdlUsers model.User) ResponeLoginUser {
// 	return ResponeLoginUser{
// 		ID:       mdlUsers.ID,
// 		Username: mdlUsers.Username,
// 		Email:    mdlUsers.Email,
// 		Name:     mdlUsers.Name,
// 		Token:    mdlUsers.Token,
// 	}
// }

func newResponseArray(mdlUsers []model.User) []ResponseUser {
	var resp []ResponseUser

	for _, value := range mdlUsers {
		resp = append(resp, newResponse(value))
	}

	return resp
}
