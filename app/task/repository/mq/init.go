package mq

import (
	"fmt"
	"go-micro-todoList/config"
	"log"

	"github.com/streadway/amqp"
)

var RabbitMq *amqp.Connection

func InitRabbitMQ() {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/", config.RabbitMQ, config.RabbitMQUser, config.RabbitMQPassword, config.RabbitHost, config.RabbitPort)
	fmt.Println("InitRabbitMQ", connString)
	conn, err := amqp.Dial(connString)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err.Error())
	}
	RabbitMq = conn
}
