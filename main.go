package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saifoelloh/minimalism_api/api"
	"github.com/saifoelloh/minimalism_api/utilities/database"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:  "GET,PUT,POST,DELETE,OPTIONS",
		ExposeHeaders: "Content-Type,Authorization,Accept",
	}))

	database.Connect()

	route := app.Group("/api/v1")
	api.RouterSetup(route)

	log.Fatal(app.Listen(":3000"))
}
