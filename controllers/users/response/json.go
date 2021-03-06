package response

import (
	"time"
	"yukevent/business/users"
)

type UserRegisterResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain users.Domain) UserRegisterResponse {
	return UserRegisterResponse{
		Message:   "Registration Success",
		ID:        domain.ID,
		Username:  domain.Username,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainLogin(domain users.Domain) UserLoginResponse {
	return UserLoginResponse{
		Message: "Login Success",
		Token:   domain.Token,
	}
}
