package main

import (
	handler "devcode-todo/handlers"
	"devcode-todo/helper"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file :", err.Error())
	}
	PORT := helper.GetEnv("PORT", "3030")

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	activity := app.Group("/activity-groups")
	{
		activity.Get("", handler.Activities)
		activity.Post("", handler.CreateActivity)
		activity.Get("/:id", handler.Activity)
		activity.Delete("/:id", handler.DeleteActivity)
		activity.Patch("/:id", handler.UpdateActivity)
	}

	todo := app.Group("/todo-items")
	{
		todo.Get("", handler.Todos)
		todo.Post("", handler.CreateTodo)
		todo.Get("/:id", handler.Todo)
		todo.Delete("/:id", handler.DeleteTodo)
		todo.Patch("/:id", handler.UpdateTodo)
	}

	app.Listen(":" + PORT)
}
