package factory

import (
	"user-service/src/cache"
	"user-service/src/publisher"
	"user-service/src/repository"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Factory struct {
	UserRepository    repository.User
	CacheRepository   cache.Cache
	RabbitMQPublisher *publisher.RabbitMQ
}

func New(db *gorm.DB, rc *redis.Client, pr *publisher.RabbitMQ) *Factory {
	return &Factory{
		UserRepository:  repository.NewUser(db),
		CacheRepository: cache.NewCache(rc),
		RabbitMQPublisher: pr,
	}
}
