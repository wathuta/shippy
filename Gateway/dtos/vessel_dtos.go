package dtos

import "time"

type CreateVesselRequest struct {
	//to do add request_id using uuid
	Message string    `json:"message"`
	Vessel  Vesseldto `json:"vessel"`
}
type GetAllVesselRequest struct {
	Message string `json:"GetAll_Vessel"`
}
type CreateVesselResponse struct {
	VesselID   int
	Created_at time.Time
}
type GetAllVesselsResponse struct {
	Vessels []Vesseldto
}
type Vesseldto struct {
	Id        int    `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"maxweight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerId   string `json:"owner_id"`
}
