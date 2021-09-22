package config

import (
	"booking-ticket/models/admins"
	"booking-ticket/models/cinemas"
	"booking-ticket/models/movies"
	"booking-ticket/models/schedules"
	"booking-ticket/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/booking_ticket?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&users.User{})
	DB.AutoMigrate(&admins.Admin{})
	DB.AutoMigrate(&movies.Movie{})
	DB.AutoMigrate(&cinemas.Cinema{})
	DB.AutoMigrate(&schedules.Schedule{})
}
