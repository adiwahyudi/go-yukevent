package events

import (
	"fmt"
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
		fmt.Println("Error mysql.go")
		return events.Domain{}, result.Error
	}

	return toDomain(ev), nil

}
