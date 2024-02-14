package auth

import (
	"Fibers-Example/model/data/request/auth"
	"github.com/gofiber/fiber/v3"
)

type LoginAPI struct {
}

func (l LoginAPI) Handler(ctx fiber.Ctx) error {
	body, err := auth.LoginRequest{}.ValidatedParse(ctx.Body())
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	return ctx.Status(200).JSON(body)
}
