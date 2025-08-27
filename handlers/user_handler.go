package handlers

import (
	"fmt"
	"lingxiaoyun/internal/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	RoleAdmin   Role = "管理员"
	RoleTeacher Role = "教师"
	RoleStudent Role = "学生"
	RoleParent  Role = "家长"
)

func ParseRole(r string) (Role, error) {
	switch Role(r) {
	case RoleAdmin, RoleTeacher, RoleStudent, RoleParent:
		return Role(r), nil
	default:
		return "", fmt.Errorf("角色无效: %s", r)
	}
}

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Role      Role      `gorm:"type:varchar(20);not null" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}

func toUserResponse(u User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "获取用户失败"})
	}

	var resp []UserResponse
	for _, u := range users {
		resp = append(resp, toUserResponse(u))
	}

	return c.JSON(resp)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "用户未找到"})
	}
	return c.JSON(toUserResponse(user))
}

func CreateUser(c *fiber.Ctx) error {
	var input CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "请求数据错误"})
	}

	role, err := ParseRole(input.Role)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "密码加密失败"})
	}

	user := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "创建用户失败"})
	}

	return c.Status(201).JSON(toUserResponse(user))
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "用户未找到"})
	}

	var input UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "请求数据错误"})
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Role != "" {
		role, err := ParseRole(input.Role)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		user.Role = role
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "密码加密失败"})
		}
		user.Password = string(hashedPassword)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "更新用户失败"})
	}

	return c.JSON(toUserResponse(user))
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&User{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "删除用户失败"})
	}
	return c.JSON(fiber.Map{"message": "用户删除成功"})
}
