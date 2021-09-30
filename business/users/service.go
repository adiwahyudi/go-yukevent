package users

import "context"

type serviceUser struct {
	userRepository Repository
}

func NewServiceUser(repoUser Repository) Service {
	return &serviceUser{
		userRepository: repoUser,
	}
}

func (serv *serviceUser) Register(ctx context.Context, domain Domain) (Domain, error) {

	result, err := serv.userRepository.Register(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

// func (serv *serviceUser) Login(ctx context.Context, email, password string) (Domain, error) {

// 	return domain, nil
// }
