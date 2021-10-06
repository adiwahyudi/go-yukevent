package events

import (
	"time"
	"yukevent/business/events"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	ID               int       `json:"id" gorm:"primary_key"`
	OrganizerID      int       `json:"organizer_id"`
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
	UpdatedAt        time.Time `json:"updated_at" form:"updated_at"`
}

func toDomain(ev Events) events.Domain {
	return events.Domain{
		ID:               ev.ID,
		OrganizerID:      ev.OrganizerID,
		Name:             ev.Name,
		Description:      ev.Description,
		Syarat_Ketentuan: ev.Syarat_Ketentuan,
		Location:         ev.Location,
		Event_Date:       ev.Event_Date,
		Event_Time:       ev.Event_Time,
		Close_Register:   ev.Close_Register,
		Capacity:         ev.Capacity,
		Poster:           ev.Poster,
		Price:            ev.Price,
		CreatedAt:        ev.CreatedAt,
		UpdatedAt:        ev.UpdatedAt,
	}
}

func fromDomain(domain events.Domain) Events {
	return Events{
		ID:               domain.ID,
		OrganizerID:      domain.OrganizerID,
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
