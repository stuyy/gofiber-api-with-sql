package handlers

import "github.com/gofiber/fiber/v2"

func GetTodosHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}

func ValidateCreateTodo(ctx *fiber.Ctx) error {
  
	return ctx.Next()
}

func CreateTodoHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(201)
}
