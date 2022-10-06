package main

import (
	"fmt"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
)

var (
	topic = "USER_CREATION"
)

type Subscriber struct{}

func main() {
	service := micro.NewService(micro.Name("shippy.email"), micro.Version("latest"))
	service.Init()

	// micro.RegisterSubscriber(USER_CREATION_TOPIC, service.Server(), new(Subscriber))
	// if err := service.Run(); err != nil {
	// 	log.Println(err)
	// }
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	Sub()
	// select {}
	service.Run()
}
func Sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[Sub] received message:", string(p.Message().Body))
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
