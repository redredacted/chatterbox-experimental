package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/RemoteENv-Team/broker/log"
)

func main() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	
	if err != nil {
		log.Logger().Fatal("Could not establish connection with RabbitMQ: " + err.Error())
	}

	log.Logger().Info("connection established")
	conn.Close()
}