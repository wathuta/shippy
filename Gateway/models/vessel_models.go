package models

type Vessel struct {
	ID        string
	Capacity  int
	MaxWeight int
	Name      string
	Available bool
	Owner     string
	CreatedAt string
}

type Specification struct {
	Capacity   int32
	Max_height int32
}
type Reponse struct {
	_Vessel Vessel
	Vessels []Vessel
	Created bool
}
