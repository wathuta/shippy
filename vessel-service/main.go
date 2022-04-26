package main

import (
	"log"
	"vessel_service/proto/vessel"
	"vessel_service/repository"
	"vessel_service/services"

	"github.com/joho/godotenv"
	micro "go-micro.dev/v4"
)
func init()  {
	if err:=godotenv.Load("");err!=nil{

	}
}

func main() {
	vessels := []*vessel.Vessel{{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500}}
	service := micro.NewService(micro.Name("vessel-service"))
	service.Init()
	repo := repository.New_vessel_repository(vessels)
	vessel_service := services.New_vessel_service(repo)
	
	if err := vessel.RegisterVesselServiceHandler(service.Server(), vessel_service); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
