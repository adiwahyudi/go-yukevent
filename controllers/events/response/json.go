package response

import (
	"time"
	"yukevent/business/events"
)

type CreateEventResponse struct {
	Message          string    `json:"message"`
	ID               int       `json:"id:"`
	Name             string    `json:"name"`
	Description      string    `json:"desc"`
	Syarat_Ketentuan string    `json:"s&k"`
	Location         string    `json:"loc"`
	Event_Date       string    `json:"event_date"`
	Event_Time       string    `json:"event_time"`
	Close_Register   string    `json:"close_regist"`
	Capacity         int       `json:"capacity"`
	Poster           string    `json:"poster"`
	Price            int       `json:"price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FromDomainCreate(domain events.Domain) CreateEventResponse {
	return CreateEventResponse{
		Message:          "Create Event Success",
		ID:               domain.ID,
		Name:             domain.Name,
		Description:      domain.Description,
		Syarat_Ketentuan: domain.Syarat_Ketentuan,
		Location:         domain.Location,
		Event_Date:       domain.Event_Date,
		Event_Time:       domain.Event_Time,
		Close_Register:   domain.Close_Register,
		Capacity:         domain.Capacity,
		Poster:           domain.Poster,
		Price:            domain.Price,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

type AllEventResponse struct {
	ID               int       `json:"id:"`
	Name             string    `json:"name"`
	Description      string    `json:"desc"`
	Syarat_Ketentuan string    `json:"s&k"`
	Location         string    `json:"loc"`
	Event_Date       string    `json:"event_date"`
	Event_Time       string    `json:"event_time"`
	Close_Register   string    `json:"close_regist"`
	Capacity         int       `json:"capacity"`
	Poster           string    `json:"poster"`
	Price            int       `json:"price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FromDomainAllEvent(domain events.Domain) AllEventResponse {
	return AllEventResponse{
		ID:               domain.ID,
		Name:             domain.Name,
		Description:      domain.Description,
		Syarat_Ketentuan: domain.Syarat_Ketentuan,
		Location:         domain.Location,
		Event_Date:       domain.Event_Date,
		Event_Time:       domain.Event_Time,
		Close_Register:   domain.Close_Register,
		Capacity:         domain.Capacity,
		Poster:           domain.Poster,
		Price:            domain.Price,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}
func FromEventListDomain(domain []events.Domain) []AllEventResponse {
	var response []AllEventResponse
	for _, value := range domain {
		response = append(response, FromDomainAllEvent(value))
	}
	return response
}
