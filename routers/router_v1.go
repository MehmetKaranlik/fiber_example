package routers

import (
	v1 "Fibers-Example/api/v1"
	"Fibers-Example/api/v1/auth"
	"github.com/gofiber/fiber/v3"
)

const (
	API_PREFIX = "/api/v1"
)

type RouterV1 struct{}

func (self RouterV1) Bind(app *fiber.App) {
	app.Get("/", v1.IndexAPI{}.Handler)
	app.Post(API_PREFIX+"/auth/login", auth.LoginAPI{}.Handler)
}
