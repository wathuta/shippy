package repository

import "log"

type VesselRepository interface {
	CreateVessel()
	GetAllVessels()
}
type vessel_Repository struct {
	l *log.Logger
}

func NewVesselRepository(l *log.Logger) VesselRepository {
	return &vessel_Repository{
		l: l,
	}
}
func (v vessel_Repository) CreateVessel() {

}
func (v vessel_Repository) GetAllVessels() {

}
