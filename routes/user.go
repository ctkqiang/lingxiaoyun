package routes

import (
	"lingxiaoyun/handlers"

	"github.com/gofiber/fiber/v2"
)

/*
UserRoutes 绑定用户相关 API

GET /users
  获取所有用户
  curl:
    curl -X GET http://localhost:3000/users

GET /users/:id
  获取单个用户
  curl:
    curl -X GET http://localhost:3000/users/1

POST /users
  创建用户
  普通用户 / 教师 / 学生 / 家长:
    curl -X POST http://localhost:3000/users \
    -H "Content-Type: application/json" \
    -d '{"name":"张三","email":"zhangsan@example.com","password":"123456","role":"学生"}'

  管理员（必须提供 auth_code = 1300177）:
    curl -X POST http://localhost:3000/users \
    -H "Content-Type: application/json" \
    -d '{"name":"李四","email":"lisi@example.com","password":"123456","role":"管理员","auth_code":"1300177"}'

PUT /users/:id
  更新用户信息（可选更新 name/email/role/password）
  curl:
    curl -X PUT http://localhost:3000/users/1 \
    -H "Content-Type: application/json" \
    -d '{"name":"新名字","password":"newpassword"}'

DELETE /users/:id
  删除用户
  curl:
    curl -X DELETE http://localhost:3000/users/1
*/

func UserRoutes(app *fiber.App) {
	users := app.Group("/users")

	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUser)
	users.Post("/", handlers.CreateUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)
}
