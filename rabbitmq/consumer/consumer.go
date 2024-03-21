package main

import (
	"fmt"
	common "gorabbitmq/rabbitmq/common"
	model "gorabbitmq/rabbitmq/model"
)

func main() {
	rmq := &model.RabbitMQ{}
	common.Start(rmq)
	common.Channel(rmq)
	Consumer(rmq)
	common.Close(rmq)
}

func Consumer(rmq *model.RabbitMQ) {
	Consume(rmq)
	BlockToUse(rmq)
}

func Consume(rmq *model.RabbitMQ) {
	rmq.Delivery, rmq.Err = rmq.Channel.Consume(
		"Testing",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	common.FailOnError(rmq.Err, "Consume Fail")
}

func BlockToUse(rmq *model.RabbitMQ) {
	forever := make(chan bool)
	go func() {
		for d := range rmq.Delivery {
			fmt.Printf("Received Messages: %s\n", d.Body)
		}
	}()
	<-forever
}
