package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://flight:yyabc123@139.198.13.12:4128/Flight")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
