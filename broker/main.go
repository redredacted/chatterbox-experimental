package main

import (
	"github.com/RemoteENv-Team/broker/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {	
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.LogStdout().Fatal("Failed to connect to RabbitMQ: " + err.Error())
		logger.FLogInfo("Failed to connect to RabbitMQ")
		logger.FLogFatal(err)
	}
	
	logger.LogStdout().Info("Connected to RabbitMQ")
	
	defer conn.Close()
}