package handler

import (
	"github.com/renaldyhidayatt/fiberEntCrud/entity"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"

	"github.com/gofiber/fiber/v2"
)

type handlerAuth struct {
	auth entity.EntityAuth
}

func NewHandlerAuth(auth entity.EntityAuth) *handlerAuth {
	return &handlerAuth{auth: auth}
}

func (h *handlerAuth) HandlerRegister(c *fiber.Ctx) error {
	var body schemas.SchemaUsers

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, errorCreate := h.auth.EntityRegister(&body)

	if errorCreate.Type == "error_register_01" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"msg": "Email already taken",
		})

	}
	if errorCreate.Type == "error_register_02" {
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"msg": "Register new user account failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Register new user account success",
	})

}
