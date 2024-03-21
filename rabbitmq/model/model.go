package model

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	ClusterToConnect string
	ConnectionIndex  int
	QueueName        string
	FileName         string
	Err              error
	Queue            amqp.Queue
	Channel          *amqp.Channel
	Notify           chan *amqp.Error
	Broker           *amqp.Connection
	Delivery         <-chan amqp.Delivery
	MsgChannel       chan map[string]interface{}
}
