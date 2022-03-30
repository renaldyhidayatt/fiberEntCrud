package main

import (
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

	routes.NewAuthRoute(client, app)

	app.Listen(":5000")

}
