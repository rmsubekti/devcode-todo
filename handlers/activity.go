package handlers

import (
	"devcode-todo/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status    string `json:"status"`
	Message   string `json:"message,omitempty"`
	Data      any    `json:"data,omitempty"`
	DeletedId *int   `json:"deletedId,omitempty"`
}

var Activities = func(c *fiber.Ctx) error {
	var activities model.Activities
	if err := activities.List(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Failed",
			Message: "List are empty",
		})
	}
	return c.JSON(&Response{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

var CreateActivity = func(c *fiber.Ctx) error {
	var activity model.Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Bad Request",
			Message: "no activity created",
		})
	}

	if err := activity.Create(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Failed",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&Response{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

var Activity = func(c *fiber.Ctx) error {
	var activity model.Activity
	id, _ := strconv.Atoi(c.Params("id"))
	if err := activity.Get(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&Response{
			Status:  "Not Found",
			Message: err.Error(),
		})
	}

	return c.JSON(&Response{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

var UpdateActivity = func(c *fiber.Ctx) error {
	var activity model.Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "Bad Request",
			Message: "no activity updated",
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))
	if err := activity.Update(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&Response{
			Status:  "Not Found",
			Message: err.Error(),
		})
	}

	return c.JSON(&Response{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}
var DeleteActivity = func(c *fiber.Ctx) error {
	var activity model.Activity
	id, _ := strconv.Atoi(c.Params("id"))
	if err := activity.Delete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&Response{
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
