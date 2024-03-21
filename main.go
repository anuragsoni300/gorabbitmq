package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rmq := &RabbitMQ{}
	Start(rmq)
	Channel(rmq)
	Queue(rmq)
	Publish(rmq)
	Close(rmq)
}

func Start(rmq *RabbitMQ) {
	rmq.Broker, rmq.Err = amqp.Dial("amqp://guest:guest@localhost:5672")
	if rmq.Err != nil {
		FailOnError(rmq.Err, "Dial Fail")
	}
}

func Channel(rmq *RabbitMQ) {
	rmq.Channel, rmq.Err = rmq.Broker.Channel()
	if rmq.Err != nil {
		FailOnError(rmq.Err, "Channel Fail")
	}
}

func Queue(rmq *RabbitMQ) {
	rmq.Queue, rmq.Err = rmq.Channel.QueueDeclare(
		"Testing",
		false,
		false,
		false,
		false,
		nil,
	)
	if rmq.Err != nil {
		FailOnError(rmq.Err, "Queue Fail")
	} else {
		fmt.Println(rmq.Queue)
	}
}

func Publish(rmq *RabbitMQ) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rmq.Err = rmq.Channel.PublishWithContext(
		ctx,
		"",
		"Testing",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("HIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII"),
		},
	)
	FailOnError(rmq.Err, "Publish Fail")
}

func Close(rmq *RabbitMQ) {
	defer rmq.Broker.Close()
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
