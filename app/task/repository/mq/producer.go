package mq

import (
	"fmt"
	"go-micro-todoList/consts"

	"github.com/streadway/amqp"
)

func SendMessage2MQ(body []byte) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	q, _ := ch.QueueDeclare(consts.RabbitMqTaskQueue, true, false, false, false, nil)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return
	}
	fmt.Println("send message to mq success")
	return
}
