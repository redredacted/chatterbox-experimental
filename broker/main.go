package main

import (
	"github.com/RemoteENv-Team/broker/tools"
	"github.com/RemoteENv-Team/broker/queueHandler"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {	
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/"); if err != nil {
		// Log to standard output and to file
		tools.LogStdout().Fatal("Failed to connect to RabbitMQ: " + err.Error())
		tools.FLogFatal(err)
	}

	queueHandler.SetupQueueHandlers(conn)

	tools.LogStdout().Info("Successfully connected to RabbitMQ!")
	defer conn.Close()
}