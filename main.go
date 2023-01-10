package main

import (
	"context"

	"github.com/renaldyhidayatt/fiberEntCrud/routes"

	"github.com/renaldyhidayatt/fiberEntCrud/config"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	client := config.Database()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello World")
	})

	context := context.Background()

	routes.NewAuthRoute(client, context, app)
	routes.NewTodoRoute(client, context, app)

	app.Listen(":5000")

}
