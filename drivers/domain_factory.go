package drivers

import (
	userDomain "yukevent/business/users"
	userDB "yukevent/drivers/databases/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}
