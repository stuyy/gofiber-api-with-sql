package main

import (
	"os"
	"web-api-sql/routes"
	"web-api-sql/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	PORT := os.Getenv("PORT")
	utils.InitDB()

	app := fiber.New()
	todos := fiber.New()

	app.Mount("/api/todos", todos)
	routes.Todos(todos)

	utils.Validator = validator.New(validator.WithRequiredStructEnabled())

	app.Listen(PORT)
}
