package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/handler"
	"github.com/renaldyhidayatt/fiberEntCrud/repository"
	"github.com/renaldyhidayatt/fiberEntCrud/services"
)

func NewTodoRoute(client *ent.Client, context context.Context, router *fiber.App) {

	repository := repository.NewTodoRepository(client, context)
	service := services.NewTodoService(repository)
	handler := handler.NewHandlerTodo(service)

	route := router.Group("/api/todo")

	route.Post("/create", handler.HandlerCreate)
	route.Get("/results", handler.HandlerResults)
}
