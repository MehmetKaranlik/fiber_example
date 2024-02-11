package v1

import "github.com/gofiber/fiber/v3"

type IndexAPI struct{}

func (i IndexAPI) Handler(ctx fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
