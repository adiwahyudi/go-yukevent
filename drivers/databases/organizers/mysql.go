package organizers

import (
	"yukevent/business"
	"yukevent/business/organizers"

	"gorm.io/gorm"
)

type MysqlOrganizersRepository struct {
	Conn *gorm.DB
}

func NewMysqlOrganizerRepository(conn *gorm.DB) organizers.Repository {
	return &MysqlOrganizersRepository{
		Conn: conn,
	}
}

func (rep *MysqlOrganizersRepository) Register(domain *organizers.Domain) (organizers.Domain, error) {

	org := fromDomain(*domain)

	result := rep.Conn.Create(&org)

	if result.Error != nil {
		return organizers.Domain{}, result.Error
	}

	return toDomain(org), nil
}

func (rep *MysqlOrganizersRepository) Login(email, password string) (organizers.Domain, error) {
	var org Organizers
	err := rep.Conn.First(&org, "email = ?", email).Error

	if err != nil {
		return organizers.Domain{}, business.ErrEmailorPass
	}

	return toDomain(org), nil
}
