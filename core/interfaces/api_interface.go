package interfaces

import "github.com/gofiber/fiber/v3"

type APIInterface interface {
	Handler(ctx fiber.Ctx) error
}
