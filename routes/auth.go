package routes

import (
	"context"

	"github.com/renaldyhidayatt/fiberEntCrud/handler"
	"github.com/renaldyhidayatt/fiberEntCrud/repository"
	"github.com/renaldyhidayatt/fiberEntCrud/services"

	"github.com/renaldyhidayatt/fiberEntCrud/ent"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRoute(client *ent.Client, context context.Context, router *fiber.App) {

	repository := repository.NewAuthRepository(client, context)
	service := services.NewAuthService(repository)
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
