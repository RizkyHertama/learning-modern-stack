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

type Postgres struct {
	DSN string
}

type Config struct {
	RabbitMQ   *rabbitmq
	CurrentApp *currentApp
	Postgres   *Postgres
}

var Conf *Config

func Load() {
	rabbitmqConf := new(rabbitmq)
	rabbitmqConf.DSN = os.Getenv("RABBITMQ_DSN")

	currentAppConf := new(currentApp)
	currentAppConf.RestfulAddress = os.Getenv("RESTFUL_ADDRESS")

	postgresConf := new(Postgres)
	postgresConf.DSN = os.Getenv("POSTGRES_DSN")

	Conf = &Config{
		RabbitMQ:   rabbitmqConf,
		CurrentApp: currentAppConf,
		Postgres:   postgresConf,
	}
}
