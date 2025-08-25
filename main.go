package main

import (
	"lingxiaoyun/internal/config"
	"lingxiaoyun/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	config.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		if c.Accepts("html") == "html" {
			return c.Render("index", fiber.Map{
				"Title": "Title",
				"Msg":   "Message lalalal",
				"Body":  "Body",
			})
		}
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	routes.UserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
