package handlers

import (
	"log"
	"net/http"

	"gateway/constants"
	"gateway/dtos"
	"gateway/services"
	"gateway/utils"

	"github.com/gin-gonic/gin"
)

type Vessel_hander struct {
	vesselService services.VesselService
	l             *log.Logger
}

func NewVessel(l *log.Logger, vesselService services.VesselService) *Vessel_hander {
	return &Vessel_hander{l: l, vesselService: vesselService}
}
func (v *Vessel_hander) CreateVessel(ctx *gin.Context) {
	req := dtos.CreateVesselRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, constants.BadRequest(ctx))
		return
	}

	if err := utils.CheckRequest(req); err != nil {
		ctx.JSON(http.StatusBadRequest, constants.BadRequestValue(ctx))
		return
	}

	response, err := v.vesselService.CreateVessel(ctx, &req, ctx.Request.Header.Get("Content-Type"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, dtos.Ok(ctx, response))
}
func (v *Vessel_hander) GetAllVessels(ctx *gin.Context) {
	v.vesselService.GetAllVessels(ctx, ctx.Request.Header.Get("Content-Type"))
}
