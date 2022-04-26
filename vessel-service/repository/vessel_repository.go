package repository

import (
	"errors"
	"log"
	"os"
	"vessel_service/model"
	"vessel_service/proto/vessel"
)

type Vessel_respository interface {
	FindAvailable(spec *model.Specification) (*vessel.Vessel, error)
}
type vessel_Repository struct {
	vessels []*vessel.Vessel
}

func New_vessel_repository(vessels []*vessel.Vessel) Vessel_respository {
	return &vessel_Repository{vessels: vessels}
}

var file, _ = os.Create("db-level.logs")

func (vr *vessel_Repository) FindAvailable(spec *model.Specification) (*vessel.Vessel, error) {
	logger := log.New(file, "\n", log.Flags())

	for _, ves := range vr.vessels {
		if spec.Capacity <= ves.Capacity && spec.MaxWeight <= ves.MaxWeight {
			logger.Println(spec, ves)
			return ves, nil
		}
	}
	return nil, errors.New("no vessel found by that spec")
}
