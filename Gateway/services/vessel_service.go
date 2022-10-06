package services

import (
	"context"
	"encoding/json"
	"gateway/constants"
	"gateway/dtos"
	"gateway/rabbitmq"
	"gateway/repository"
	"log"
)

type vesselService struct {
	VesselRepository repository.VesselRepository

	rabbit *rabbitmq.VesselQueue
}

type VesselService interface {
	CreateVessel(ctx context.Context, req *dtos.CreateVesselRequest, contentType string) (*dtos.CreateVesselResponse, error)
	GetAllVessels(ctx context.Context, contentType string)
}

func NewVesselService(VesselRepository repository.VesselRepository, rabbit *rabbitmq.VesselQueue) VesselService {
	return &vesselService{
		VesselRepository: VesselRepository,
		rabbit:           rabbit,
	}
}
func (v *vesselService) CreateVessel(ctx context.Context, req *dtos.CreateVesselRequest, contentType string) (*dtos.CreateVesselResponse, error) {
	log.Println("Create vessel service")
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err = v.rabbit.Publish([]byte(body), contentType, constants.CREATE_VESSEL_QUEUE, constants.CREATE_VESSEL_ROUTING_KEY); err != nil {
		return nil, err
	}
	return &dtos.CreateVesselResponse{VesselID: req.Vessel.Id}, nil
}
func (v *vesselService) GetAllVessels(ctx context.Context, contentType string) {
	log.Println("Getting all vessels")
	req := dtos.GetAllVesselRequest{Message: "GET ALL VESSELS"}
	body, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return
	}
	if err := v.rabbit.Publish([]byte(body), contentType, constants.FIND_ALL_VESSEL_QUEUE, constants.FIND_ALL_VESSEL_ROUTING_KEY); err != nil {
		log.Println(err)
		return
	}
}
