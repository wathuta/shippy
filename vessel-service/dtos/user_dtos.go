package dtos

import "vessel_service/model"

type Create_user_request struct {
	Id        string
	Capacity  int32
	MaxWeight int32
	Name      string
	Available bool
	OwnerId   string
}
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
