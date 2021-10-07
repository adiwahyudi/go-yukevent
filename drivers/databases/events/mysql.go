package events

import (
	"yukevent/business"
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

func (rep *MysqlEventRepository) Update(orgID int, evID int, domain *events.Domain) (events.Domain, error) {

	eventUpdate := fromDomain(*domain)

	eventUpdate.ID = evID

	result := rep.Conn.Where("organizer_id = ?", orgID).Where("id = ?", evID).Updates(&eventUpdate)

	if result.Error != nil {
		return events.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(eventUpdate), nil
}

func (rep *MysqlEventRepository) Delete(orgID int, id int) (string, error) {
	rec := Events{}

	find := rep.Conn.Where("organizer_id = ?", orgID).Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Event has been delete", nil

}

func (rep *MysqlEventRepository) MyEventByOrganizer(orgID int) ([]events.Domain, error) {
	var event []Events

	result := rep.Conn.Where("organizer_id = ?", orgID).Find(&event)

	if result.Error != nil {
		return []events.Domain{}, result.Error
	}

	return toDomainList(event), nil
}

func (rep *MysqlEventRepository) EventByID(id int) (events.Domain, error) {

	var event Events

	result := rep.Conn.Where("id = ?", id).First(&event)

	if result.Error != nil {
		return events.Domain{}, result.Error
	}

	return toDomain(event), nil
}
func (rep *MysqlEventRepository) EventByIdOrganizer(orgzID int) ([]events.Domain, error) {

	var event []Events

	result := rep.Conn.Where("organizer_id = ?", orgzID).Find(&event)

	if result.Error != nil {
		return []events.Domain{}, result.Error
	}

	return toDomainList(event), nil
}
