package routes

import (
	"lingxiaoyun/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/api/users")

	user.Get("/", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
}
