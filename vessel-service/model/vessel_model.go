package model

import "time"

type Vessel struct {
	Id        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"max_weight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerId   string `json:"owner_id"`

	Created_at time.Time `json:"created_at"`
}

type Specification struct {
	Capacity  int32
	MaxWeight int32
}
