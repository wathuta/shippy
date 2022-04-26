package services

import (
	"context"
	"log"
	"vessel_service/mapper"
	"vessel_service/proto/vessel"
	"vessel_service/repository"
)

type VesselServiceHandler interface {
	FindAvailable(ctx context.Context, req *vessel.Specification, res *vessel.Response) error
	Create(ctx context.Context, in *vessel.Vessel, out *vessel.Response) error
}
type vesselServiceHandler struct {
	repo repository.Vessel_respository
}

func New_vessel_service(repo repository.Vessel_respository) VesselServiceHandler {
	return &vesselServiceHandler{repo: repo}
}
//Finds available vessel
func (v *vesselServiceHandler) FindAvailable(ctx context.Context, req *vessel.Specification, res *vessel.Response) error {
	//Converts spec to spec model
	spec := mapper.Marshal_Spec(req)
	
	newVessel, err := v.repo.FindAvailable(spec)
	if err != nil {
		log.Println(err)
		return err
	}
	res.Vessel = newVessel
	return nil
}
func (v *vesselServiceHandler) Create(ctx context.Context, in *vessel.Vessel, out *vessel.Response) error {
	return nil
}
