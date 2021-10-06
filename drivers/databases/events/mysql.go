package events

import (
	"yukevent/business/events"

	"gorm.io/gorm"
)

type MysqlEventRepository struct {
	Conn *gorm.DB
}

func NewMysqlEventRepository(conn *gorm.DB) events.Repository {
	return &MysqlEventRepository{
		Conn: conn,
	}
}

func (rep *MysqlEventRepository) Create(orgID int, domain *events.Domain) (events.Domain, error) {

	ev := fromDomain(*domain)

	ev.OrganizerID = orgID

	result := rep.Conn.Create(&ev)

	if result.Error != nil {
		return events.Domain{}, result.Error
	}

	return toDomain(ev), nil

}

func (rep *MysqlEventRepository) AllEvent() ([]events.Domain, error) {

	var event []Events

	result := rep.Conn.Find(&event)

	if result.Error != nil {
		return []events.Domain{}, result.Error
	}

	return toDomainList(event), nil

}
