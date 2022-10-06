package dtos

import "vessel_service/model"

type Find_user_request struct {
	Capacity  int32
	MaxWeight int32
}
type Create_user_response struct {
	Created bool
}
type FindAll_users_response struct {
	Vessels model.Vessel
}
type Vessel_dtos struct {
	Id        int    `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"maxweight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerId   string `json:"owner_id"`
}
type CreateVesselRequest struct {
	//to do add request_id using uuid
	Message string      `json:"message"`
	Vessel  Vessel_dtos `json:"vessel"`
}
type GetAllVesselRequest struct {
	Message string `json:"message"`
}
