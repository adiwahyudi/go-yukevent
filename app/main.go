package main

import (
	"log"

	_routes "yukevent/app/routes"

	_userService "yukevent/business/users"
	_userController "yukevent/controllers/users"
	_userRepo "yukevent/drivers/databases/users"

	_organizerService "yukevent/business/organizers"
	_organizerController "yukevent/controllers/organizers"
	_organizerRepo "yukevent/drivers/databases/organizers"

	_dbDriver "yukevent/drivers/mysql"

	_driverFactory "yukevent/drivers"

	_middleware "yukevent/app/middleware"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_organizerRepo.Organizers{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: int64(viper.GetInt(`jwt.expired`)),
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewServiceUser(userRepo, 10, &configJWT)
	userCtrl := _userController.NewControllerUser(userService)

	organizerRepo := _driverFactory.NewOrganizerRepository(db)
	organizerService := _organizerService.NewServiceOrganizer(organizerRepo, 10, &configJWT)
	organizerCtrl := _organizerController.NewControllerOrganizer(organizerService)

	routesInit := _routes.ControllerList{
		JWTMiddleware:       configJWT.Init(),
		UserController:      *userCtrl,
		OrganizerController: *organizerCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
