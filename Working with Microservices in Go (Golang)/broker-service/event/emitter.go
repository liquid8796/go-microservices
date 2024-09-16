package event

import (
	_ "github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	return declareExchange(channel)
}
