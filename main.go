package main

import (
	"context"
	"log"

	"github.com/renaldyhidayatt/fiberEntCrud/routes"

	"github.com/renaldyhidayatt/fiberEntCrud/config"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	context := context.Background()
	client, err := config.Database(context)

	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello World")
	})

	routes.NewAuthRoute(client, context, app)
	routes.NewTodoRoute(client, context, app)

	app.Listen(":5000")

}
