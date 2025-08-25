package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.Accepts("application/xml")
		c.Accepts("text/html")

		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
