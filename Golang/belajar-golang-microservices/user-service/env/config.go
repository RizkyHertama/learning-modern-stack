package env

import (
	"os"
)

type rabbitmq struct {
	DSN string
}

type currentApp struct {
	RestfulAddress string
}

type Config struct {
	RabbitMQ *rabbitmq
	CurrentApp *currentApp
}

var Conf *Config

func Load() {
	rabbitmqConf := new(rabbitmq)
	rabbitmqConf.DSN = os.Getenv("RABBITMQ_DSN")

	currentAppConf := new(currentApp)
	currentAppConf.RestfulAddress = os.Getenv("RESTFUL_ADDRESS")

	Conf = &Config{
		RabbitMQ: rabbitmqConf,
	}
}
