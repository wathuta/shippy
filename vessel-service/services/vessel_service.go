package services

import (
	"context"
	"vessel_service/mapper"
	"vessel_service/proto/vessel"
	"vessel_service/repository"

	"github.com/sirupsen/logrus"
)

type VesselServiceHandler interface {
	FindAvailableVessel(ctx context.Context, req *vessel.Specification, res *vessel.Response) error
	CreateVessel(ctx context.Context, in *vessel.Vessel, out *vessel.Response) error
}
type vesselServiceHandler struct {
	repo  repository.Vessel_respository
	loger *logrus.Entry
}

func New_vessel_service(repo repository.Vessel_respository, loger *logrus.Entry) VesselServiceHandler {

	return &vesselServiceHandler{repo: repo, loger: loger}
}

//Finds available vessel
func (v *vesselServiceHandler) FindAvailableVessel(ctx context.Context, req *vessel.Specification, res *vessel.Response) error {
	//Converts spec to spec model
	spec := mapper.Marshal_Spec(req)
	// spec.Capacity = 0
	newVessel, err := v.repo.FindBySpec(ctx, spec)
	if err != nil {
		v.loger.Error("unable to find file due to ", err)
		return err
	}
	v.loger.Info("found user", newVessel)
	res.Vessel = mapper.Unmarshal_Vessel(newVessel)
	return nil
}
func (v *vesselServiceHandler) CreateVessel(ctx context.Context, in *vessel.Vessel, out *vessel.Response) error {
	vessel := mapper.Marshal_Vessel(in)
	created, err := v.repo.Create(ctx, vessel)
	if err != nil {
		v.loger.Error("Unable to create vessel", err)
		return err
	}
	v.loger.Info("Created", created)
	out.Created = true
	return nil
}
