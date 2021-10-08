package organizers

import (
	"time"
	"yukevent/business/organizers"

	"gorm.io/gorm"
)

type Organizers struct {
	gorm.Model
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	Name         string
	Phone_Number string
	Photo        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func toDomain(org Organizers) organizers.Domain {
	return organizers.Domain{
		ID:           org.ID,
		Username:     org.Username,
		Email:        org.Email,
		Password:     org.Password,
		Name:         org.Name,
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
		Phone_Number: domain.Phone_Number,
		Photo:        domain.Photo,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
