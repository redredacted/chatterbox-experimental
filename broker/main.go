package main

import (
	"github.com/RemoteENv-Team/broker/logger"
	"github.com/RemoteENv-Team/broker/queueHandler"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {	
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/"); if err != nil {
		// Log to standard output and to file
		logger.LogStdout().Fatal("Failed to connect to RabbitMQ: " + err.Error())
		logger.FLogFatal(err)
	}

	queueHandler.SetupQueueHandlers(conn)

	logger.LogStdout().Info("Successfully connected to RabbitMQ!")
	defer conn.Close()
}