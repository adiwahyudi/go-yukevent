package organizers

import (
	"time"
	"yukevent/app/middleware"
	"yukevent/business"
	"yukevent/helpers/encrypt"
)

type serviceOrganizer struct {
	organizerRepository Repository
	contextTimeout      time.Duration
	jwtAuth             *middleware.ConfigJWT
}

func NewServiceOrganizer(repoOrganizer Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) Service {
	return &serviceOrganizer{
		organizerRepository: repoOrganizer,
		contextTimeout:      timeout,
		jwtAuth:             jwtauth,
	}
}

func (serv *serviceOrganizer) Register(domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.organizerRepository.Register(domain)

	if result == (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return result, nil
}

func (serv *serviceOrganizer) Login(email, password string) (Domain, error) {

	result, err := serv.organizerRepository.Login(email, password)

	if err != nil {
		return Domain{}, business.ErrEmailorPass
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return Domain{}, business.ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "organizer")

	return result, nil
}
