package main

import (
	"yukevent/config"
	"yukevent/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")
}
