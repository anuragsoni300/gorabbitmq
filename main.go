package main

import (
	common "gorabbitmq/rabbitmq/common"
	model "gorabbitmq/rabbitmq/model"
	p "gorabbitmq/rabbitmq/producer"
)

func main() {
	rmq := &model.RabbitMQ{}
	common.Start(rmq)
	common.Channel(rmq)
	p.Producer(rmq)
	common.Close(rmq)
}
