package request

import (
	"yukevent/business/organizers"
)

type Organizer struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone_Number string `json:"phone_number"`
	Photo        string `json:"photo"`
}

func (req *Organizer) ToDomain() *organizers.Domain {
	return &organizers.Domain{
		Username:     req.Username,
		Password:     req.Password,
		Email:        req.Email,
		Name:         req.Name,
		Phone_Number: req.Phone_Number,
		Photo:        req.Photo,
	}
}

type OrganzierLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
