package env

import (
	"os"
)

type rabbitmq struct {
	DSN string
}

type Config struct {
	RabbitMQ *rabbitmq
}

var Conf *Config

func Load() {
	rabbitmqConf := new(rabbitmq)
	rabbitmqConf.DSN = os.Getenv("RABBITMQ_DSN")

	Conf = &Config{
		RabbitMQ: rabbitmqConf,
	}
}
