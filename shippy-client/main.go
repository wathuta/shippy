package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/wathuta/shippy-client/proto/consignment"
	user "github.com/wathuta/shippy-client/proto/user"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/metadata"
)

const (
	defaultFilename = "consignment.json"
	address         = "localhost:50051"
)

func parsefile(filename string) (*pb.Consignment, error) {
	consignment := &pb.Consignment{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(file, &consignment)
	return consignment, nil
}
func main() {
	service := micro.NewService(micro.Name("shippy.client"))
	service.Init()

	client := pb.NewShippingService("shippy.consignment", service.Client())
	userclient := user.NewUserService("shippy.user", service.Client())
	file := defaultFilename
	log.Println(os.Args)
	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecing file and token."))
	}
	file = os.Args[1]
	token := os.Args[2]

	var consignment, err = parsefile(file)
	if err != nil {
		checkError(err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{"token": token})

	resp, err := client.CreateConsignment(ctx, consignment)
	if err.Error()[0:5] == "token" {
		_token, err := userclient.AuthUser(context.Background(), &user.User{Email: "wathutabrian@gmail.com", Password: "CChangez13115!@"})
		if err != nil {
			checkError(err)
		}
		ctx := metadata.NewContext(context.Background(), map[string]string{"token": _token.Token})
		resp, err = client.CreateConsignment(ctx, consignment)
		checkError(err)
	} else {
		checkError(err)
	}
	log.Println("Create:", resp.Created)

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

//var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImVtYWlsIjoid2F0aHV0YWJyaWFuQGdtYWlsLmNvbSIsInBhc3N3b3JkIjoiQ0NoYW5nZXoxMzExNSFAIn0sImV4cCI6MTUwMDAsImlzcyI6InNoaXBweV9zZXJ2aWNlIn0.srIFF0usezxiQhxYIm5_3MAmIPc13Z9RQ7sXP0xuEQw"
