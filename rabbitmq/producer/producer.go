package producer

import (
	"context"
	"fmt"
	common "gorabbitmq/rabbitmq/common"
	model "gorabbitmq/rabbitmq/model"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// func main() {
// 	rmq := &model.RabbitMQ{}
// 	common.Start(rmq)
// 	common.Channel(rmq)
// 	Producer(rmq)
// 	common.Close(rmq)
// }

func Producer(rmq *model.RabbitMQ) {
	Queue(rmq)
	Publish(rmq)
}

func Queue(rmq *model.RabbitMQ) {
	rmq.Queue, rmq.Err = rmq.Channel.QueueDeclare(
		"Testing",
		false,
		false,
		false,
		false,
		nil,
	)
	if rmq.Err != nil {
		common.FailOnError(rmq.Err, "Queue Fail")
	} else {
		fmt.Println(rmq.Queue)
	}
}

func Publish(rmq *model.RabbitMQ) {
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
			Body:        []byte(uuid.New().String()),
		},
	)
	common.FailOnError(rmq.Err, "Publish Fail")
}
