package main

import (
	//"fmt"
	"github.com/Fabricio2210/gofiber/elastic"
	"github.com/Fabricio2210/gofiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	elastic.ConnectElastic()

	app := fiber.New()

	app.Use(cors.New())
	router.DefaultRouter(app, "DSP")
	router.DefaultRouter(app, "DDM")
	router.DefaultRouter(app, "RAW")
	router.DefaultRouter(app, "POP")
	router.DefaultRouter(app, "SHINKO")
	router.DefaultRouter(app, "reacts")
	app.Listen(":3000")
}
