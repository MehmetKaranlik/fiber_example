package auth

import (
	"Fibers-Example/model/data/request/auth"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type LoginAPI struct {
}

func (l LoginAPI) Handler(ctx fiber.Ctx) error {
	body, err := auth.LoginRequest{}.ValidatedParse(ctx.Body())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(body)
}
