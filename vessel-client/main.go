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
	specArr := []vessel.Specification{
		{Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000},
		{Capacity: 500, MaxWeight: 55000}, {Capacity: 00, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 500},
		{Capacity: 500, MaxWeight: 55000}, {Capacity: 500, MaxWeight: 55000}, {Capacity: 5300, MaxWeight: 55000}, {Capacity: 50320, MaxWeight: 550542300}, {Capacity: 50000000, MaxWeight: 55000000}, {Capacity: 500, MaxWeight: 55000000},
	}
	// spec := vessel.Specification{Capacity: 500, MaxWeight: 55000}

	newVesselService := vessel.NewVesselService("vessel-service", service.Client())
	for _, spec := range specArr {
		vesselresponse, err := newVesselService.FindAvailable(context.Background(), &spec)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Println(vesselresponse)
	}
}
