package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
	"github.com/renaldyhidayatt/fiberEntCrud/service"
)

type handlerTodo struct {
	todo *service.ServiceTodo
}

func NewHandlerTodo(todo *service.ServiceTodo) *handlerTodo {
	return &handlerTodo{todo: todo}
}

func (h *handlerTodo) HandlerCreate(c *fiber.Ctx) error {

	var body schemas.SchemaTodo
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res, err := h.todo.EntityCreate(&body)

	if err.Type == "error_create_01" {
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"msg": "Create new todo failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Create new todo success",
		"data":  res,
	})
}

func (h *handlerTodo) HandlerResults(c *fiber.Ctx) error {

	res, err := h.todo.EntityResults()

	if err.Type == "error_results_01" {
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"msg": "Get todo failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Get todo success",
		"data":  &res,
	})

}
