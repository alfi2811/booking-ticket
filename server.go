package main

import (
	"booking-ticket/config"
	"booking-ticket/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":1323"))
}
