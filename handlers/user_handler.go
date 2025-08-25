package handlers

import (
	"lingxiaoyun/internal/config"
	"lingxiaoyun/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var user []models.User

	config.DB.Find(&user)
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	return c.JSON("Create User")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON("Update User")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.JSON("Delete User")
}
