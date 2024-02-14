package main

import (
	"Fibers-Example/routers"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log"
	"os"
)

var (
	ipAddr = os.Getenv("IP")
	port   = os.Getenv("PORT")
)

func main() {
	app := fiber.New()
	routers.RouterV1{}.Bind(app)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", ipAddr, port)))
}
