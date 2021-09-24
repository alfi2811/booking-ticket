package main

import (
	"booking-ticket/app/routes"
	_adminUsecase "booking-ticket/business/admins"
	_cinemaUsecase "booking-ticket/business/cinemas"
	_movieUsecase "booking-ticket/business/movies"
	_scheduleUsecase "booking-ticket/business/schedules"
	_userUsecase "booking-ticket/business/users"

	_adminController "booking-ticket/controllers/admins"
	_cinemaController "booking-ticket/controllers/cinemas"
	_movieController "booking-ticket/controllers/movies"
	_scheduleController "booking-ticket/controllers/schedules"
	_userController "booking-ticket/controllers/users"
	_userdb "booking-ticket/drivers/databases/users"

	_mysqlDriver "booking-ticket/drivers/mysql"
	"time"

	_adminRepository "booking-ticket/drivers/databases/admins"
	_cinemaRepository "booking-ticket/drivers/databases/cinemas"
	_movieRepository "booking-ticket/drivers/databases/movies"
	_scheduleRepository "booking-ticket/drivers/databases/schedules"
	_userRepository "booking-ticket/drivers/databases/users"

	_middleware "booking-ticket/app/middlewares"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userdb.Users{})
	db.AutoMigrate(&_movieRepository.Movies{})
	db.AutoMigrate(&_cinemaRepository.Cinemas{})
	db.AutoMigrate(&_scheduleRepository.Schedules{})
}

func main() {
	// init koneksi databse
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	configJWTAdmin := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwtAdmin.secret`),
		ExpiresDuration: viper.GetInt(`jwtAdmin.expired`),
	}

	adminRepository := _adminRepository.NewMysqlAdminRepository(Conn)
	adminUseCase := _adminUsecase.NewAdminUsecase(adminRepository, timeoutContext, configJWTAdmin)
	adminController := _adminController.NewAdminController(adminUseCase)

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUseCase)

	movieRepository := _movieRepository.NewMysqlMovieRepository(Conn)
	movieUseCase := _movieUsecase.NewMovieUsecase(movieRepository, timeoutContext)
	movieController := _movieController.NewMovieController(movieUseCase)

	cinemaRepository := _cinemaRepository.NewMysqlCinemaRepository(Conn)
	cinemaUseCase := _cinemaUsecase.NewCinemaUsecase(cinemaRepository, timeoutContext)
	cinemaController := _cinemaController.NewCinemaController(cinemaUseCase)

	scheduleRepository := _scheduleRepository.NewMysqlScheduleRepository(Conn)
	scheduleUseCase := _scheduleUsecase.NewScheduleUsecase(scheduleRepository, timeoutContext)
	scheduleController := _scheduleController.NewScheduleController(scheduleUseCase)

	routesInit := routes.ControllerList{
		JwtConfig:          configJWT.Init(),
		JwtConfigAdmin:     configJWTAdmin.Init(),
		UserController:     *userController,
		AdminController:    *adminController,
		MovieController:    *movieController,
		CinemaController:   *cinemaController,
		ScheduleController: *scheduleController,
	}

	apiV1 := e.Group("api/v1/")
	routesInit.RouteUsers(apiV1)
	routesInit.RouteAdmins(apiV1)
	routesInit.RouteMovies(apiV1)
	routesInit.RouteCinemas(apiV1)
	routesInit.RouteSchedule(apiV1)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
