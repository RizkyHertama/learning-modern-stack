package handler

import (
	"user-service/src/common/dto/request"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (h *Auth) Register(c *fiber.Ctx) error {
	req := new(request.Register)

	if err := c.BodyParser(req); err != nil {
		return err
	}
}
