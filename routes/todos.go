package routes

import (
	"web-api-sql/handlers"

	"github.com/gofiber/fiber/v2"
)

func Todos(app *fiber.App) {
	app.Get("/", handlers.GetTodosHandler)
	app.Post("/", handlers.ValidateCreateTodo, handlers.CreateTodoHandler)
}
