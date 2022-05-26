package routes

import (
	"github.com/renaldyhidayatt/fiberEntCrud/service"

	"github.com/renaldyhidayatt/fiberEntCrud/handler"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRoute(client *ent.Client, router *fiber.App) {

	service := service.NewServiceAuth(client)
	handler := handler.NewHandlerAuth(service)

	route := router.Group("/api/auth")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"error":   false,
			"message": "Ready",
		})
	})

	route.Post("/register", handler.HandlerRegister)
	route.Post("/login", handler.HandlerLogin)

}
