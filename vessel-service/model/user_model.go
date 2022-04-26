package model

type Vessel struct {
	Id        string
	Capacity  int32
	MaxWeight int32
	Name      string
	Available bool
	OwnerId   string
}

type Specification struct {
	Capacity  int32
	MaxWeight int32
}
