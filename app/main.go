package main

import (
	"booking-ticket/app/routes"
	_cinemaUsecase "booking-ticket/business/cinemas"
	_movieUsecase "booking-ticket/business/movies"
	_userUsecase "booking-ticket/business/users"
	_cinemaController "booking-ticket/controllers/cinemas"
	_movieController "booking-ticket/controllers/movies"
	_userController "booking-ticket/controllers/users"
	_userdb "booking-ticket/drivers/databases/users"
	_mysqlDriver "booking-ticket/drivers/mysql"
	"time"

	_cinemaRepository "booking-ticket/drivers/databases/cinemas"
	_movieRepository "booking-ticket/drivers/databases/movies"
	_userRepository "booking-ticket/drivers/databases/users"

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

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	movieRepository := _movieRepository.NewMysqlMovieRepository(Conn)
	movieUseCase := _movieUsecase.NewMovieUsecase(movieRepository, timeoutContext)
	movieController := _movieController.NewMovieController(movieUseCase)

	cinemaRepository := _cinemaRepository.NewMysqlCinemaRepository(Conn)
	cinemaUseCase := _cinemaUsecase.NewCinemaUsecase(cinemaRepository, timeoutContext)
	cinemaController := _cinemaController.NewCinemaController(cinemaUseCase)

	routesInit := routes.ControllerList{
		UserController:   *userController,
		MovieController:  *movieController,
		CinemaController: *cinemaController,
	}

	routesInit.RouteUsers(e)
	routesInit.RouteMovies(e)
	routesInit.RouteCinemas(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
