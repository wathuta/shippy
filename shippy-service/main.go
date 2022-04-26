package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wathuta/shippy-service/proto/consignment"
	"github.com/wathuta/shippy-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wathuta/shippy-service/configs"
	"github.com/wathuta/shippy-service/repository"
	"github.com/wathuta/shippy-service/services"
	micro "go-micro.dev/v4"
)

var (
	client *mongo.Client
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Loading env file error %v", err)
		log.Println("searching for environment variables")
	}
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		log.Fatal("Empty uri field")
	}
	var err error
	client, err = configs.CreateConn(context.Background(), uri, 0)
	if err != nil {
		log.Fatalf("Error unable to connect to mongo DB: %v ", err)
	}
}

func main() {
	service := micro.NewService(micro.Name("consignment-service"))
	service.Init()

	//to do add a collection and connect to mongodb

	collection := client.Database("shippy").Collection("consignments")
	defer client.Disconnect(context.Background())
	//
	repo := repository.New_shippy_repository(collection)
	vess := vessel.NewVesselService("vessel-service", service.Client())
	if vess == nil {
		log.Fatal(vess)
	}

	/**/
	spec := vessel.Specification{Capacity: 500, MaxWeight: 55000}
	vesselresponse, err := vess.FindAvailable(context.Background(), &spec)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(vesselresponse)
	// log.Fatal(vess).
	/**/

	shippy_service := services.New_shippy_service(repo, vess)

	if err := consignment.RegisterShippingServiceHandler(service.Server(), shippy_service); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
