package drivers

import (
	eventDomain "yukevent/business/events"
	organizerDomain "yukevent/business/organizers"
	userDomain "yukevent/business/users"

	eventDB "yukevent/drivers/databases/events"
	organizerDB "yukevent/drivers/databases/organizers"
	userDB "yukevent/drivers/databases/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}

func NewOrganizerRepository(conn *gorm.DB) organizerDomain.Repository {
	return organizerDB.NewMysqlOrganizerRepository(conn)
}

func NewEventRepository(conn *gorm.DB) eventDomain.Repository {
	return eventDB.NewMysqlEventRepository(conn)
}
