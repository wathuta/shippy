package main

import (
	"context"
	"log"
	"os"
	"vessel_service/configs"
	"vessel_service/constants"
	"vessel_service/proto/vessel"
	"vessel_service/rabbit"
	"vessel_service/repository"
	"vessel_service/services"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	micro "go-micro.dev/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func init() {
	if err := godotenv.Load("vars.env"); err != nil {
		log.Printf("Loading env file error %v", err)
		log.Println("searching for environment variables ...")
	}
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		log.Fatal("Empty uri field")
	}
	var err error

	client, err = configs.CreateConn(context.Background(), uri, 1)
	if err != nil {
		log.Fatalf("Error unable to connect to mongo DB: %v ", err)
	}
}

func main() {
	service := micro.NewService(micro.Name("shippy.vessel"))
	service.Init()

	//connecting to a specific collection in the mongo cluster
	collection := client.Database("shippy").Collection("vessels")
	defer client.Disconnect(context.Background())

	loggers := logrus.New()
	repo := repository.New_vessel_repository(GenLogger(*loggers)(), collection)
	go ConnectRabbitMq(GenLogger(*loggers)(), repo)
	vessel_service := services.New_vessel_service(repo, GenLogger(*loggers)())

	if err := vessel.RegisterVesselServiceHandler(service.Server(), vessel_service); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func GenLogger(l logrus.Logger) func() *logrus.Entry {
	// filePath, _ := filepath.Abs("../logs/" + "logs.log")
	// Create db if it doesn't exist
	file, err := os.Create("logs.log")
	if err != nil {
		log.Fatal(err)
	}
	return func() *logrus.Entry {
		contentLoger := l.WithFields(logrus.Fields{"service": "Vessel_service"})
		contentLoger.Logger.SetReportCaller(true)
		contentLoger.Logger.Out = file
		contentLoger.Logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
		return contentLoger
	}
}

func ConnectRabbitMq(l *logrus.Entry, repo repository.Vessel_respository) {
	consumer := rabbit.NewVesselQueue(configs.NewRabbitConfigs(), l, repo)
	forever := make(chan bool)
	go func() {
		if err := consumer.ConsumeFromCreateQueue(constants.CREATE_VESSEL_QUEUE, constants.CREATE_VESSEL_CONSUMER_TAG, constants.CREATE_VESSEL_ROUTING_KEY); err != nil {
			l.Fatalln("Unable to consume from Find_all queue", err)
		}
	}()
	go func() {
		if err := consumer.ConsumeFromFindAll(constants.FIND_ALL_VESSEL_QUEUE, constants.FIND_ALL_VESSEL_CONSUMER_TAG, constants.FIND_ALL_VESSEL_ROUTING_KEY); err != nil {
			l.Fatalln("Unable to consume from Find_all queue", err)
		}
	}()
	<-forever
}
