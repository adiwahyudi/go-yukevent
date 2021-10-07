package events

import "time"

type Domain struct {
	ID               int
	OrganizerID      int
	Name             string
	Description      string
	Syarat_Ketentuan string
	Location         string
	Event_Date       string
	Event_Time       string
	Close_Register   string
	Capacity         int
	Poster           string
	Price            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Service interface {
	AllEvent() ([]Domain, error)
	Create(orgID int, domain *Domain) (Domain, error)
	Update(orgID int, evID int, domain *Domain) (Domain, error)
	Delete(orgID, id int) (string, error)
	MyEventByOrganizer(orgID int) ([]Domain, error)
	EventByID(id int) (Domain, error)
	EventByIdOrganizer(orgzID int) ([]Domain, error)
}

type Repository interface {
	AllEvent() ([]Domain, error)
	Create(orgID int, domain *Domain) (Domain, error)
	Update(orgID int, evID int, domain *Domain) (Domain, error)
	Delete(orgID, id int) (string, error)
	MyEventByOrganizer(orgID int) ([]Domain, error)
	EventByID(id int) (Domain, error)
	EventByIdOrganizer(orgzID int) ([]Domain, error)
}
