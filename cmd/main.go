package main

import (
	"Fibers-Example/routers"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
	"os"
)

var (
	ipAddr = os.Getenv("IP")
	port   = os.Getenv("PORT")
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	routers.RouterV1{}.Bind(app)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", ipAddr, port)))
}
