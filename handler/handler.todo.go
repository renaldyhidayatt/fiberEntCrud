package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
	"github.com/renaldyhidayatt/fiberEntCrud/services"
)

type handlerTodo struct {
	todo services.TodoService
}

func NewHandlerTodo(todo services.TodoService) *handlerTodo {
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

	res, err := h.todo.Create(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Create new todo success",
		"data":  res,
	})
}

func (h *handlerTodo) HandlerResults(c *fiber.Ctx) error {

	res, err := h.todo.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Get todo success",
		"data":  &res,
	})

}

func (h *handlerTodo) HandlerFindById(c *fiber.Ctx) error {
	id := c.Params("id")

	idparam, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res, err := h.todo.FindById(idparam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Get todo success",
		"data":  &res,
	})
}

func (h *handlerTodo) UpdateById(c *fiber.Ctx) error {
	id := c.Params("id")

	idparam, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var body schemas.SchemaTodo
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res, err := h.todo.UpdateById(idparam, body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "update todo success",
		"data":  res,
	})
}

func (h *handlerTodo) DeleteById(c *fiber.Ctx) error {

	id := c.Params("id")

	idparam, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	res, err := h.todo.DeleteById(idparam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "delete todo success",
		"data":  res,
	})
}
