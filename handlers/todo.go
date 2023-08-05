package handlers

import (
	"devcode-todo/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var Todos = func(c *fiber.Ctx) error {
	var todos model.Todos
	activityID := c.Query("activity_group_id")

	if len(activityID) > 0 {
		id, _ := strconv.Atoi(activityID)

		if err := todos.ListByActivity(id); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&Response{
				Status:  "Failed",
				Message: "Data empty",
			})
		}

		return c.JSON(&Response{
			Status:  "Success",
			Message: "Success",
			Data:    todos,
		})
	}

	if err := todos.List(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Failed",
			Message: "Data empty",
		})
	}
	return c.JSON(&Response{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

var CreateTodo = func(c *fiber.Ctx) error {
	var todo model.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Failed",
			Message: "Data empty",
		})
	}

	if err := todo.Create(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Failed",
			Message: err.Error(),
		})
	}

	return c.JSON(&Response{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}

var Todo = func(c *fiber.Ctx) error {
	var todo model.Todo
	id, _ := strconv.Atoi(c.Params("id"))
	if err := todo.Get(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Not Found",
			Message: err.Error(),
		})
	}

	return c.JSON(&Response{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}

var UpdateTodo = func(c *fiber.Ctx) error {
	var todo model.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Failed",
			Message: "Data empty",
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))
	if err := todo.Update(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Not Found",
			Message: err.Error(),
		})
	}

	return c.JSON(&Response{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}
var DeleteTodo = func(c *fiber.Ctx) error {
	var todo model.Todo
	id, _ := strconv.Atoi(c.Params("id"))
	if err := todo.Delete(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Not Found",
			Message: err.Error(),
		})
	}

	return c.JSON(&Response{
		Status:    "Success",
		Message:   "Success",
		DeletedId: &id,
	})
}
