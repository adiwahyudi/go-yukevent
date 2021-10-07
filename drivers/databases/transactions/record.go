package transactions

import (
	"time"
	"yukevent/business/transactions"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	UserID    int
	EventID   int
	Status    bool
	Uniq_Code string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(tr Transactions) transactions.Domain {
	return transactions.Domain{
		ID:        tr.ID,
		UserID:    tr.UserID,
		EventID:   tr.EventID,
		Status:    tr.Status,
		Uniq_code: tr.Uniq_Code,
		CreatedAt: tr.CreatedAt,
		UpdatedAt: tr.UpdatedAt,
	}
}

func fromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		ID:        domain.ID,
		UserID:    domain.UserID,
		EventID:   domain.EventID,
		Status:    domain.Status,
		Uniq_Code: domain.Uniq_code,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainList(data []Transactions) []transactions.Domain {
	result := []transactions.Domain{}

	for _, trans := range data {
		result = append(result, toDomain(trans))
	}
	return result
}
