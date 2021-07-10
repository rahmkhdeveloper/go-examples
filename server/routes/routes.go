package routes

import (
	"github.com/gofiber/fiber"
	"mydomain.com/companyname/gofiber-example/handlers"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", handlers.Register)
	app.Post("/user", handlers.GetUser)
}
