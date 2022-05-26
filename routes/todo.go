package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/handler"
	"github.com/renaldyhidayatt/fiberEntCrud/repository"
	"github.com/renaldyhidayatt/fiberEntCrud/service"
)

func NewTodoRoute(client *ent.Client, router *fiber.App) {
	repository := repository.NewRepositoryTodo(client)
	service := service.NewServiceTodo(repository)
	handler := handler.NewHandlerTodo(service)

	route := router.Group("/api/todo")

	route.Post("/create", handler.HandlerCreate)
	route.Get("/results", handler.HandlerResults)
}
