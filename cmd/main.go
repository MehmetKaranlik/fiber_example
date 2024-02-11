package main

import (
	"Fibers-Example/routers"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"os"
)

var (
	ip_addr = os.Getenv("IP")
)

func main() {
	app := fiber.New()
	routers.RouterV1{}.Bind(app)
	port := os.Getenv("PORT")
	err := app.Listen(fmt.Sprintf("%s:%s", ip_addr, port))
	if err != nil {
		return
	}
}
