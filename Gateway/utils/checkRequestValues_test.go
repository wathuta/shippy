package utils

import (
	"gateway/dtos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRequest(t *testing.T) {
	req := dtos.CreateVesselRequest{
		"Create_vessel",
		dtos.Vesseldto{Id: 0, Capacity: 200, MaxWeight: 50000, Name: "my vessel", Available: true},
	}
	assert.Equal(t, CheckRequest(req), nil)
}

//go test -coverprofile cover.out
//go tool cover -html=cover.out
