package main

import (
	"context"
	"log"
	"os"
	"vessel-client/proto/vessel"

	micro "go-micro.dev/v4"
)

func main() {
	service := micro.NewService(micro.Name("vessel-client"))
	service.Init()
	spec := vessel.Specification{Capacity: 500, MaxWeight: 55000}
	newVesselService := vessel.NewVesselService("vessel-service", service.Client())

	vesselresponse, err := newVesselService.FindAvailable(context.Background(), &spec)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(vesselresponse)
}
