package response

import (
	"time"
	"yukevent/business/organizers"
)

type OrganizerRegisterRespone struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain organizers.Domain) OrganizerRegisterRespone {
	return OrganizerRegisterRespone{
		Message:   "Organizer Registration Success",
		ID:        domain.ID,
		Username:  domain.Username,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type OrganizerLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainLogin(domain organizers.Domain) OrganizerLoginResponse {
	return OrganizerLoginResponse{
		Message: "Organizer Login Success",
		Token:   domain.Token,
	}
}
