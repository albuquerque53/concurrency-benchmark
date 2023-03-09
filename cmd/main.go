package main

import (
	"concurrencybenchmark/internal/application/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/default", handler.Default)
	app.Post("/concurrent", handler.Concurrent)

	app.Listen(":3000")
}
