package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/wathuta/shippy-client/proto/consignment"
	micro "go-micro.dev/v4"
)

const (
	defaultFilename = "consignment.json"
	address         = "localhost:50051"
)

func main() {
	service := micro.NewService(micro.Name("shippy-client"))
	service.Init()

	client := pb.NewShippingService("consignment-service", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	var consignment *pb.Consignment
	bytesSlice, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("we are unable to read from this file", err)
		log.Fatal(err)
	}

	err = json.Unmarshal(bytesSlice, &consignment)
	if err != nil {
		log.Println("we are unable to unmarshal the consignment", err)
		log.Fatal(err)
	}

	resp, err := client.CreateConsignment(context.Background(), consignment)
	checkError(err)

	log.Println("Create:", resp.Created)
	log.Println(resp)
	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, consignment := range resp.Consignments {
		log.Println(consignment.Containers)
	}
}
func checkError(err error) {
	if err != nil {
		log.Println("we are here because of ", err)
		log.Fatal(err)
	}
}
