package drivers

import (
	adminDomain "yukevent/business/admins"
	eventDomain "yukevent/business/events"
	organizerDomain "yukevent/business/organizers"
	transDomain "yukevent/business/transactions"
	userDomain "yukevent/business/users"

	adminDB "yukevent/drivers/databases/admins"
	eventDB "yukevent/drivers/databases/events"
	organizerDB "yukevent/drivers/databases/organizers"
	transDB "yukevent/drivers/databases/transactions"
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

func NewTransRepository(conn *gorm.DB) transDomain.Repository {
	return transDB.NewMysqlTransRepository(conn)
}

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}

