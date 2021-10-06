package events

import "fmt"

type serviceEvent struct {
	eventRepository Repository
}

func NewServiceEvent(repoEvent Repository) Service {
	return &serviceEvent{
		eventRepository: repoEvent,
	}
}

// func (serv *serviceEvent) AllEvent() ([]Domain, error) {

// 	result, err := serv.eventRepository.AllEvent()

// 	if err != nil {
// 		return []Domain{}, err
// 	}

// 	return result, nil
// }

func (serv *serviceEvent) Create(orgID int, domain *Domain) (Domain, error) {

	result, err := serv.eventRepository.Create(orgID, domain)

	if err != nil {
		fmt.Println("error di service")
		return Domain{}, err
	}

	return result, nil
}

// func (serv *serviceEvent) Update(domain *Domain) (Domain, error) {

// 	result, err := serv.eventRepository.Update(domain)

// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return result, nil
// }

// func (serv *serviceEvent) Delete(id int) error {

// 	err := serv.eventRepository.Delete(id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
