package handler

import (
	"encoding/base64"
	"time"
	dto "user-service/src/common/dto/request"
	"user-service/src/common/dto/response"
	"user-service/src/factory"
	"user-service/src/restful/service"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	service service.Auth
}

func NewAuth(f *factory.Factory) *Auth {
	return &Auth{
		service: service.NewAuth(f),
	}
}

func (h *Auth) Register(c *fiber.Ctx) error {
	req := new(dto.Register)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	email, err := h.service.Register(c, req)
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{
		Name: "pending_register",
		Value: base64.RawStdEncoding.EncodeToString([]byte(email)),
		HTTPOnly: true,
		Path: "/api/v1/auth/register/verify",
		Expires: time.Now().Add(30 * time.Minute),
	})

	return c.Status(200).JSON(response.Common{Message: "request success, please check your email for otp"})
}
