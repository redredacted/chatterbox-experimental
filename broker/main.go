package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	
	if err != nil {
		Logger().Fatal("could not establish connection with RabbitMQ:" + err.Error())
	}

	Logger().Info("connection established")
	conn.Close()
}