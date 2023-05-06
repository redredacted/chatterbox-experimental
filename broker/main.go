package main

import (
	"fmt"

	"github.com/RemoteENv-Team/broker/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {	
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Info("Failed to connect to RabbitMQ")
		logger.Fatal(err)
	}
	
	fmt.Println("Connected to RabbitMQ")
	
	defer conn.Close()
}