package rabbitmq

import (
	"gateway/configs"
	"gateway/retry"
	"log"

	"github.com/streadway/amqp"
)

//to switch to env
const url = "amqp://guest:guest@localhost:5672/"

type VesselQueue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	cfg        *configs.Config
}

func NewRabbitMqConn(cfg *configs.Config) *VesselQueue {
	var Conn *amqp.Connection
	log.Println("Connection to RabbitMQ..")
	err := retry.Do(func(attempt int) (retry bool, err error) {
		//to do use configs
		Conn, err = amqp.Dial(url)
		return attempt < 4, err
	})
	if err != nil {
		log.Fatalf("unable to connect to RabbitMq server %v", err)
	}
	myChan, err := Conn.Channel()
	if err != nil {
		log.Fatalf("unable to connect to RabbitMq server %v", err)
	}

	return &VesselQueue{Connection: Conn, cfg: cfg, Channel: myChan}
}

func (c *VesselQueue) PublisherExchangeQueue(queueName, routingKey string) error {
	log.Println("Seting up rabbitMQ routing")
	err := c.Channel.ExchangeDeclare(c.cfg.Exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}
	queue, err := c.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	log.Println(queue.Name)
	err = c.Channel.QueueBind(queue.Name, routingKey, c.cfg.Exchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}
func (c *VesselQueue) Publish(body []byte, contentType, queueName, routingKey string) error {
	err := c.PublisherExchangeQueue(queueName, routingKey)
	if err != nil {
		return err
	}
	log.Println("Publishing the content to", queueName, "-queue")
	if err := c.Channel.Publish(c.cfg.Exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  contentType,
		Body:         body,
	}); err != nil {
		return err
	}
	return nil
}
