package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"rabbit/helpers"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.FailOnError(err, "Failed to connect to rabbit")

	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed to open channel")

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	helpers.FailOnError(err, "Failed creating a queue")

	body := "Hello world in the other side"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	helpers.FailOnError(err, "Failed on publish")

	fmt.Println("Sent", body)
	defer ch.Close()
}
