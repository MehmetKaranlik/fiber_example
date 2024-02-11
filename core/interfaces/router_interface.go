package interfaces

import "github.com/gofiber/fiber/v3"

type RouterInterface interface {
	Bind(app *fiber.App)
}
