//go:build wireinject
// +build wireinject

package providers

import (
	"gateway/handlers"
	"gateway/rabbitmq"
	"gateway/repository"
	"gateway/services"
	"log"

	"github.com/google/wire"
)

func ProvideVesselHandler(l *log.Logger, vesselQueue *rabbitmq.VesselQueue) *handlers.Vessel_hander {
	wire.Build(
		services.NewVesselService,
		handlers.NewVessel,
		repository.NewVesselRepository,
	)
	return &handlers.Vessel_hander{}
}

//go run github.com/google/wire/cmd/wire
