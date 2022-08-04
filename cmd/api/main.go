package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	Sender   string `json:"sender" binding:"required"`
	Receiver string `json:"receiver" binding:"required"`
	Content  string `json:"message" binding:"required"`
}

func sendMessage(ch *amqp.Channel, msg *Message) {
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Printf("failed to declare a queue RABBITMQ: %v", err)
	}

	// msg
	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("hello, world"),
		})
		if err != nil {
			log.Printf("failed to publish a message RABBITMQ: %v", err)
		}
	log.Printf(" [x] Sent %s\n", body)
}


func main() {

	conn, err := amqp.Dial("amqp://user:password@localhost:7001/")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to create channel in RABBITMQ: %v", err)
	}

	defer ch.Close()

	r := gin.Default()

	r.POST("/message", func(c *gin.Context) {
		var msg Message
		if err := c.BindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		
		sendMessage(ch, &msg)
		c.JSON(200, "successful")
	})

	r.Run()
}
