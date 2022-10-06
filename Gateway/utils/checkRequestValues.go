package utils

import (
	"fmt"
	"gateway/dtos"
)

func CheckRequest(req dtos.CreateVesselRequest) error {
	if req.Vessel.Id > 0 || req.Vessel.Capacity < 0 || req.Vessel.MaxWeight < req.Vessel.Capacity {
		return fmt.Errorf(" bad request values")
	}
	return nil
}
