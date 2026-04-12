package service

import (
	"time"

	"user-service/src/cache"
	dto "user-service/src/common/dto/request"
	customerrors "user-service/src/common/errors"
	"user-service/src/common/util"
	"user-service/src/factory"
	"user-service/src/publisher"
	"user-service/src/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	Register(c *fiber.Ctx, data *dto.Register) (string, error)
}

type authImpl struct {
	userRepository  repository.User
	cacheRepository cache.Cache
	RabbitMQPublisher *publisher.RabbitMQ
}


func NewAuth(f *factory.Factory) Auth {
	return &authImpl{
		userRepository:  f.UserRepository,
		cacheRepository: f.CacheRepository,
		RabbitMQPublisher: f.RabbitMQPublisher,
	}
}

func (s *authImpl) Register(c *fiber.Ctx, data *dto.Register) (string, error) {
	user, err := s.userRepository.Find(c.Context(), "email = ?", []any{data.Email})
	if err != nil {
		return "", err
	}

	if user != nil {
		return "", &customerrors.Response{HttpCode: 409, Message: "Email already exists"}
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	data.Password = string(bcryptPassword)

	if err := s.cacheRepository.Set(c.Context(), "register:"+data.Email, data, 20*time.Minute); err != nil {
		return "", err
	}

	otp := util.GenerateOTP()

	sendOTPReq := dto.SendOTP{
		Email: data.Email,
		Otp:   otp,
	}

	if err := s.cacheRepository.Set(c.Context(), "otp:"+data.Email, sendOTPReq, 20*time.Minute); err != nil {
		return "", err
	}

	if err := s.RabbitMQPublisher.Publish("user", "otp", sendOTPReq); err != nil {
		return "", err
	}

	return data.Email, nil
}
