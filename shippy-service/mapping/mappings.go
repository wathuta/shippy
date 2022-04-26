package mapping

import (
	"strconv"

	"github.com/wathuta/shippy-service/models"
	"github.com/wathuta/shippy-service/proto/consignment"
	"go-micro.dev/v4/util/log"
)

//MarshalContainerCollection maps slices of generated protobuf code to slices of model instances.
func MarshalContainerCollection(containers []*consignment.Container) []*models.Container {
	collection := make([]*models.Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}
	return collection
}
func UnmarshalContainerCollection(containers []*models.Container) []*consignment.Container {
	collection := make([]*consignment.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}
	return collection
}

func MarshalContainer(container *consignment.Container) *models.Container {
	return &models.Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
		Origin:     container.Origin,
	}
}
func UnmarshalContainer(container *models.Container) *consignment.Container {
	return &consignment.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
		Origin:     container.Origin,
	}
}
func UnmarshalConsgnmentCollection(consignments []*models.Consignment) []*consignment.Consignment {
	collection := make([]*consignment.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnmarshalConsignment(consignment))
	}
	return collection
}
func UnmarshalConsignment(_consignment *models.Consignment) *consignment.Consignment {
	return &consignment.Consignment{
		Id:          _consignment.ID,
		Description: _consignment.Description,
		Containers:  UnmarshalContainerCollection(_consignment.Containers),
		VesselId:    _consignment.VesselID,
		Weight:      string(_consignment.Weight),
	}
}
func MarshalConsignmentCollection(consignments []*consignment.Consignment) []*models.Consignment {
	collection := make([]*models.Consignment, 0)
	for _, _consignment := range consignments {
		collection = append(collection, MarshalConsignment(_consignment))
	}
	return collection
}

func MarshalConsignment(_consignment *consignment.Consignment) *models.Consignment {
	return &models.Consignment{
		ID:          _consignment.Id,
		Description: _consignment.Description,
		Containers:  MarshalContainerCollection(_consignment.Containers),
		Weight:      _consignment.Weight,
		VesselID:    _consignment.VesselId,
	}
}
func StringConverter(s string) int32 {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Error(err)
	}
	return int32(num)
}
