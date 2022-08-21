package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"rabbit/helpers"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.FailOnError(err, "Connection failed")

	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed on create a channel")

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)

	helpers.FailOnError(err, "Failed to create a queue")

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	helpers.FailOnError(err, "Failed when trying to get messages")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Println("Received message:", d.Body)
		}
	}()

	log.Println("Waiting for messages")
	<-forever
}
