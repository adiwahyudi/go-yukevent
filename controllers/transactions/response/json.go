package response

import "yukevent/business/transactions"

type CreateTransResponse struct {
	Message   string `json:"message"`
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	EventID   int    `json:"event_id"`
	Status    bool   `json:"status"`
	Uniq_code string `json:"uniq_code"`
}

func FromDomainCreate(domain transactions.Domain) CreateTransResponse {
	return CreateTransResponse{
		Message:   "Transactions Success",
		ID:        domain.ID,
		UserID:    domain.UserID,
		EventID:   domain.EventID,
		Status:    domain.Status,
		Uniq_code: domain.Uniq_code,
	}
}

type TransResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	EventID   int    `json:"event_id"`
	Status    bool   `json:"status"`
	Uniq_code string `json:"uniq_code"`
}

func FromDomainAllTrans(domain transactions.Domain) TransResponse {
	return TransResponse{
		ID:        domain.ID,
		UserID:    domain.UserID,
		EventID:   domain.EventID,
		Status:    domain.Status,
		Uniq_code: domain.Uniq_code,
	}
}

func FromTransListDomain(domain []transactions.Domain) []TransResponse {
	var response []TransResponse
	for _, value := range domain {
		response = append(response, FromDomainAllTrans(value))
	}
	return response
}
