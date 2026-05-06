package publisher

import (
	"encoding/json"
	"log"

	"github.com/Sahadat20/distributed-video-processing-pipeline.git/producer/internal/domain"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   *amqp091.Queue
}

func NewRabbitMQPublisher() *RabbitMQPublisher {
	conn, err := amqp091.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatal("Faild to connect rabbitmq", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Faild to open channel", err)
	}

	q, err := ch.QueueDeclare(
		"video_processing_task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Faild to declare queue", err)
	}

	return &RabbitMQPublisher{
		conn:    conn,
		channel: ch,
		queue:   &q,
	}
}

func (p *RabbitMQPublisher) Publish(task domain.Task) error {
	body, _ := json.Marshal(task)

	return p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
