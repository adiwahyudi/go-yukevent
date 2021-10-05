package users

import (
	"yukevent/business"
	"yukevent/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Register(domain *users.Domain) (users.Domain, error) {

	var findDuplicate Users

	user := fromDomain(*domain)

	err := rep.Conn.First(&findDuplicate, "email = ?", user.Email).Error

	if err != nil {
		return users.Domain{}, nil
	}

	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return toDomain(user), nil
}

func (rep *MysqlUserRepository) Login(email, password string) (users.Domain, error) {
	var user Users
	err := rep.Conn.First(&user, "email = ?", email).Error

	if err != nil {
		return users.Domain{}, business.ErrEmailorPass
	}

	return toDomain(user), nil
}
