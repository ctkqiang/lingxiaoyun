package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"lingxiaoyun/internal/config"
	"lingxiaoyun/internal/models"
	"lingxiaoyun/routes"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移所有模型（v1 ERP 架构）
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Teacher{},
		&models.Parent{},
		&models.Student{},
		&models.Class{},
		&models.Attendance{},
		&models.Leave{},
		&models.Mark{},
		&models.Course{},
		&models.Material{},
		&models.Announcement{},
		&models.Comment{},
	)

	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	// HTML/JSON 测试路由
	app.Get("/", func(c *fiber.Ctx) error {
		if c.Accepts("html") == "html" {
			return c.Render("index", fiber.Map{
				"Title": "School ERP",
				"Msg":   "Welcome 灵儿 🚀",
				"Body":  "Your ERP system is alive!",
			})
		}
		return c.JSON(fiber.Map{"message": "Hello, ERP World!"})
	})

	app.Get("/auth/login", func(c *fiber.Ctx) error {

		if c.Accepts("html") == "html" {
			return c.Render("login", fiber.Map{
				"Title": "Login",
			})
		}

		return c.JSON(fiber.Map{"message": "Login"})
	})

	routes.UserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
