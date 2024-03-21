package common

import (
	"log"
	model "gorabbitmq/rabbitmq/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Start(rmq *model.RabbitMQ) {
	rmq.Broker, rmq.Err = amqp.Dial("amqp://guest:guest@localhost:5672")
	if rmq.Err != nil {
		FailOnError(rmq.Err, "Dial Fail")
	}
}

func Channel(rmq *model.RabbitMQ) {
	rmq.Channel, rmq.Err = rmq.Broker.Channel()
	if rmq.Err != nil {
		FailOnError(rmq.Err, "Channel Fail")
	}
}

func Close(rmq *model.RabbitMQ) {
	defer rmq.Broker.Close()
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}