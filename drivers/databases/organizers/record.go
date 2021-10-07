package organizers

import (
	"time"
	"yukevent/business/organizers"

	"gorm.io/gorm"
)

type Organizers struct {
	gorm.Model
	ID           int       `json:"id" form:"id" gorm:"primary_key"`
	Username     string    `json:"username" form:"username" gorm:"unique"`
	Email        string    `json:"email" form:"email" gorm:"unique"`
	Password     string    `json:"password" form:"password"`
	Name         string    `json:"name" form:"name"`
	Dob          string    `json:"dob" form:"dob"`
	Phone_Number string    `json:"phone_number" form:"phone_number"`
	Photo        string    `json:"photo" form:"photo"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}

func toDomain(org Organizers) organizers.Domain {
	return organizers.Domain{
		ID:           org.ID,
		Username:     org.Username,
		Email:        org.Email,
		Password:     org.Password,
		Name:         org.Name,
		Dob:          org.Dob,
		Phone_Number: org.Phone_Number,
		Photo:        org.Photo,
		CreatedAt:    org.CreatedAt,
		UpdatedAt:    org.UpdatedAt,
	}
}

func fromDomain(domain organizers.Domain) Organizers {
	return Organizers{
		ID:           domain.ID,
		Username:     domain.Username,
		Email:        domain.Email,
		Password:     domain.Password,
		Name:         domain.Name,
		Dob:          domain.Dob,
		Phone_Number: domain.Phone_Number,
		Photo:        domain.Photo,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
