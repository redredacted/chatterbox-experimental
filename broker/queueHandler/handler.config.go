package queueHandler

import (
	//"github.com/RemoteENv-Team/broker/logger"
	//amqp "github.com/rabbitmq/amqp091-go"
)

/*
func SetupQueueHandlers(conn *amqp.Connection) {
	// Queue handlers
	ch, err := conn.Channel(); if err != nil {
		logger.LogStdout().Fatal("Failed to open a channel: " + err.Error())
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"consistency", // name
		true,		  // durable
		false,		  // delete when unused
		false,		  // exclusive
		false,		  // no-wait
		nil,		  // arguments
	); if err != nil {
		logger.LogStdout().Fatal("Failed to declare a queue: " + err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",		// consumer
		true,	// auto-ack
		false,	// exclusive
		false,	// no-local
		false,	// no-wait
		nil,	// args
	); if err != nil {
		logger.LogStdout().Fatal("Failed to register a consumer: " + err.Error())
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			logger.LogStdout().Info("Received a message: " + string(d.Body))
		}
	}()

	logger.LogStdout().Info("Waiting for messages. To exit press CTRL+C")
	<-forever
}*/