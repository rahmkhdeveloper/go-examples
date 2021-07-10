package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"mydomain.com/companyname/gofiber-example/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept"},
	}))

	routes.SetUpRoutes(app)
	app.Listen("localhost:8000")
}
