package request

import (
	"yukevent/business/transactions"
)

type Transactions struct {
	EventID int `json:"event_id"`
}

func (req *Transactions) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		EventID: req.EventID,
	}
}
