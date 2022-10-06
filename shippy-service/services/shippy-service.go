package services

import (
	"context"
	"fmt"
	"log"

	"github.com/wathuta/shippy-service/mapping"
	"github.com/wathuta/shippy-service/proto/consignment"
	"github.com/wathuta/shippy-service/proto/vessel"
	"github.com/wathuta/shippy-service/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type Shippy_service struct {
	repo repository.Shippy_repository
	vess vessel.VesselService
}
type shipping_Service interface {
	CreateConsignment(ctx context.Context, in *consignment.Consignment, out *consignment.Response) error
	GetConsignments(ctx context.Context, in *consignment.GetRequest, out *consignment.Response) error
}

func New_shippy_service(repo repository.Shippy_repository, vess vessel.VesselService) shipping_Service {
	return &Shippy_service{repo: repo, vess: vess}
}
func (s *Shippy_service) CreateConsignment(ctx context.Context, in *consignment.Consignment, out *consignment.Response) error {
	vessel_response, err := s.vess.FindAvailable(ctx, &vessel.Specification{
		Capacity:  int32(len(in.Containers)),
		MaxWeight: mapping.StringConverter(in.Weight),
	})
	if err != nil {
		return err
	}

	in.VesselId = vessel_response.Vessel.Id
	_consignment := mapping.MarshalConsignment(in)

	insert_id, err := s.repo.CreateConsignment(ctx, _consignment)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error when creating consignment in the repository %v", err)
	}
	log.Println(insert_id)

	out.Consignments = append(out.Consignments, in)
	out.Created = true
	return nil
}
func (s *Shippy_service) GetConsignments(ctx context.Context, in *consignment.GetRequest, out *consignment.Response) error {
	cons, err := s.repo.GetConsignments(ctx, bson.D{})
	if err != nil {
		return fmt.Errorf("getting consignment  %v", err)
	}
	// log.Println(cons)
	out.Consignments = mapping.UnmarshalConsgnmentCollection(cons)
	return nil
}
