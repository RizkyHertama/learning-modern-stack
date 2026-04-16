package router

import (
	"user-service/src/restful/handler"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App, h *handler.Auth) {
	app.Post("/api/v1/auth/register", h.Register)
}
