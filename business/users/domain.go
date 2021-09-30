package users

import (
	"context"
	"time"
)

type Domain struct {
	Id           int
	Username     string
	Email        string
	Password     string
	Name         string
	Dob          string
	Phone_Number string
	Photo        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	// Login(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	// Login(ctx context.Context, email, password string) (Domain, error)
}
