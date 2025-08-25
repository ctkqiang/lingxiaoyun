package main

import (
	"lingxiaoyun/internal/config"
	"lingxiaoyun/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.Accepts("application/xml")
		c.Accepts("text/html")

		return c.SendString("Hello, World!")
	})

	routes.UserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
