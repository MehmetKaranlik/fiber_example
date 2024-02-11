package main

import (
	"Fibers-Example/routers"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	routers.RouterV1{}.Bind(app)
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
