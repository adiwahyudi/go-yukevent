package drivers

import (
	userDomain "yukevent/business/users"
	organizerDomain "yukevent/business/organizers"

	userDB "yukevent/drivers/databases/users"
	organizerDB "yukevent/drivers/databases/organizers"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}

func NewOrganizerRepository(conn *gorm.DB) organizerDomain.Repository {
	return organizerDB.NewMysqlOrganizerRepository(conn)
}
