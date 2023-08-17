package main

import (
	"context"
	utils "github.com/Powehi-cs/Go-RabbitMQ"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	utils.FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := utils.BodyFrom(os.Args)
	err = ch.PublishWithContext(
		ctx,
		"logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
