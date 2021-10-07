package events

import (
	"time"
	"yukevent/business/events"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	ID               int `gorm:"primary_key"`
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

func toDomainUpdate(ev Events) events.Domain {
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
func toDomainList(data []Events) []events.Domain {
	result := []events.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
