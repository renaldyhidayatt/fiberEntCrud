package handler

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/renaldyhidayatt/fiberEntCrud/pkg"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
	"github.com/renaldyhidayatt/fiberEntCrud/services"

	"github.com/gofiber/fiber/v2"
)

type handlerAuth struct {
	auth services.AuthService
}

func NewHandlerAuth(auth services.AuthService) *handlerAuth {
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

	row, err := h.auth.Register(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  true,
			"msg":    err,
			"status": fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    "Register new user account success",
		"data":   row,
		"status": fiber.StatusCreated,
	})

}

func (h *handlerAuth) HandlerLogin(c *fiber.Ctx) error {
	var body schemas.SchemaUsers
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res, err := h.auth.Login(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = res.Email
	claims["user_id"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(pkg.GodotEnv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login susccses",
		"data":    t,
	})
}
