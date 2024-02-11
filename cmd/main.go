package main

import (
	"Fibers-Example/routers"
	"github.com/gofiber/fiber/v3"
	"os"
)

func main() {
	app := fiber.New()
	routers.RouterV1{}.Bind(app)
	port := os.Getenv("PORT")
	err := app.Listen("0.0.0.0:" + port)
	if err != nil {
		return
	}
}
