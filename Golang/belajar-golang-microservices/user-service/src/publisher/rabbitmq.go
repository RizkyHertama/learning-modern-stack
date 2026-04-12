package publisher

import (
	"encoding/json"
	"sync"
	"time"

	"user-service/env"
	"user-service/src/common/log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn         *amqp.Connection
	chann        *amqp.Channel
	closeChannel chan *amqp.Error // ← bukan chan *amqp.Channel
	mutex        sync.Mutex
}

func newRabbitMQ() *RabbitMQ {
	conn, err := amqp.DialConfig(env.Conf.RabbitMQ.DSN, amqp.Config{ // ← pakai env.Conf langsung
		Heartbeat: 10 * time.Second,
	})
	if err != nil {
		log.Logger.Fatal(err.Error())
	}

	closeChann := conn.NotifyClose(make(chan *amqp.Error, 1))
	chann := setupChannel(conn)

	return &RabbitMQ{
		conn:         conn,
		chann:        chann,
		closeChannel: closeChann,
	}
}

func (r *RabbitMQ) reconnect() {
	if r.conn != nil || !r.conn.IsClosed() {
		return
	}

	for {
		conn, err := amqp.DialConfig(env.Conf.RabbitMQ.DSN, amqp.Config{
			Heartbeat: 10 * time.Second,
		})

		if err != nil {
			log.Logger.Error(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		r.conn = conn
		r.chann = setupChannel(conn)
		r.closeChannel = conn.NotifyClose(make(chan *amqp.Error, 1))

		log.Logger.Info("RabbitMQ connected")
		break
	}
}

func (r *RabbitMQ) Publish(exchange, key string, msg any) error { // ← tambah parameter exchange & key
	if r.conn == nil || r.conn.IsClosed() {
		r.reconnect()
	}
	
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	publishing := amqp.Publishing{ // ← assign ke variabel
		ContentType: "application/json",
		Body:        jsonData,
	}

	if err := r.chann.Publish(exchange, key, true, false, publishing); err != nil {
		return err
	}

	return nil // ← tambah return nil
}

func (r *RabbitMQ) Close() {
	if err := r.chann.Close(); err != nil {
		log.Logger.Fatal(err.Error())
	}

	if err := r.conn.Close(); err != nil {
		log.Logger.Fatal(err.Error())
	}
}

func setupChannel(conn *amqp.Connection) *amqp.Channel {
	chann, err := conn.Channel()
	if err != nil {
		log.Logger.Fatal(err.Error())
	}

	if err := chann.ExchangeDeclare("user", "direct", true, false, false, false, nil); err != nil {
		log.Logger.Fatal(err.Error())
	}

	return chann
}
