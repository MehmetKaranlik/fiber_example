package routers

import (
	v1 "Fibers-Example/api/v1"
	"github.com/gofiber/fiber/v3"
)

const (
	API_PREFIX = "/api/v1"
)

type RouterV1 struct{}

func (self RouterV1) Bind(app *fiber.App) {
	app.Get("/", v1.IndexAPI{}.Handler)
}
