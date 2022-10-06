package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
	"vessel_service/configs"
	"vessel_service/constants"
	"vessel_service/dtos"
	"vessel_service/model"
	"vessel_service/repository"
	"vessel_service/retry"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type VesselQueue struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	cfg     *configs.Config
	loger   *logrus.Entry

	repo repository.Vessel_respository
}

const url = "amqp://guest:guest@localhost:5672/"

func NewVesselQueue(cfg *configs.Config, loger *logrus.Entry, repo repository.Vessel_respository) *VesselQueue {
	var Conn *amqp.Connection
	log.Println("Connecting to RabbitMQ...")
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

	return &VesselQueue{
		cfg:     cfg,
		Conn:    Conn,
		Channel: myChan,
		loger:   loger,
		repo:    repo,
	}
}
func (v *VesselQueue) CreateConsumer(queName, routingKey string) error {
	log.Println("Seting up rabbitMQ routing ...")
	err := v.Channel.ExchangeDeclare(v.cfg.Exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}
	queue, err := v.Channel.QueueDeclare(queName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = v.Channel.QueueBind(queue.Name, routingKey, v.cfg.Exchange, false, nil)
	if err != nil {
		return err
	}
	if err = v.Channel.Qos(1, 0, false); err != nil {
		return err
	}
	return nil
}
func (v *VesselQueue) WorkerPool(ctx context.Context, message <-chan amqp.Delivery, returnChan chan interface{}) error {
	var err error
	for delivery := range message {
		v.loger.Printf("Process Deliveried deliveryTag %v", delivery.ConsumerTag)

		switch delivery.ConsumerTag {

		case constants.CREATE_VESSEL_CONSUMER_TAG:
			var newVessel dtos.CreateVesselRequest
			err = json.Unmarshal(delivery.Body, &newVessel)
			if err := v.deliveryAck(err, delivery); err != nil {
				return err
			}
			created, err := v.repo.Create(ctx, &model.Vessel{
				Name:       newVessel.Vessel.Name,
				Capacity:   newVessel.Vessel.Capacity,
				MaxWeight:  newVessel.Vessel.MaxWeight,
				OwnerId:    newVessel.Vessel.OwnerId,
				Created_at: time.Now().Local(),
				Available:  newVessel.Vessel.Available,
			})
			if err != nil {
				return fmt.Errorf("failed to create vessel with error %v", err)
			}
			go func() { returnChan <- created }()
			log.Println(created)
		case constants.FIND_ALL_VESSEL_CONSUMER_TAG:
			var GetAllVesselRequest *dtos.GetAllVesselRequest
			err = json.Unmarshal(delivery.Body, &GetAllVesselRequest)
			if err := v.deliveryAck(err, delivery); err != nil {
				return err
			}
			findAll, err := v.repo.FindAll(ctx)
			if err != nil {
				return fmt.Errorf("failed to FindAll vessels with error %v", err)
			}
			go func() { returnChan <- findAll }()
			// log.Println(findAll)
		}

		log.Println(string(delivery.Body))
	}
	return nil
}
func (v *VesselQueue) deliveryAck(err error, delivery amqp.Delivery) error {
	if err != nil {
		err = delivery.Reject(true)
		if err != nil {
			return fmt.Errorf("err failed to requeue: delivery.reject: %v", err)
		}
		return fmt.Errorf("failed to process delivery %v", err)
	} else {
		err = delivery.Ack(false)
		if err != nil {
			return fmt.Errorf("failed to acknowledge delivery %v", err)
		}
	}
	return nil
}

//ConsumeFromCreateQueue consumes a create vessel event
func (v *VesselQueue) ConsumeFromCreateQueue(queueName, ConsumerTag, routingKey string) error {
	if err := v.CreateConsumer(queueName, routingKey); err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()
	delivery, err := v.Channel.Consume(
		queueName,
		ConsumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	errChan := make(chan error)
	resultsChan := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := v.WorkerPool(ctx, delivery, resultsChan)
		if err != nil {
			v.loger.Error(err)
			errChan <- err
		}
		log.Println(<-resultsChan)
		errChan <- nil
	}()
	wg.Wait()
	return <-errChan
}

func (v *VesselQueue) ConsumeFromFindAll(queueName, ConsumerTag, routingKey string) error {
	if err := v.CreateConsumer(queueName, routingKey); err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()
	deliveryChan, err := v.Channel.Consume(
		queueName,
		ConsumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	errChan := make(chan error)
	var wg sync.WaitGroup
	resultsChan := make(chan interface{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := v.WorkerPool(ctx, deliveryChan, resultsChan)
		if err != nil {
			errChan <- err
		}
		errChan <- nil
		v.loger.Println("findall", <-resultsChan)
	}()
	wg.Wait()
	return <-errChan
}
