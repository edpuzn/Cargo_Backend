package main

import (
	"Cargo_Dash/database"
	"Cargo_Dash/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {

	database.Connect()

	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8080",
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
