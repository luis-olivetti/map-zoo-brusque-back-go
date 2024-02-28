package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/model"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/request"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/service"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/helper"
	resp "github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/helper"
)

type MarkerHandler interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type markerHandler struct {
	*Handler
	markerService service.MarkerService
}

func (m *markerHandler) Create(ctx *gin.Context) {
	var markerRequest request.MarkerRequest
	if err := ctx.ShouldBindJSON(&markerRequest); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, "Bad Request", nil)
		return
	}

	if err := helper.ValidateRequiredFields(markerRequest); err != nil {
		m.logger.Logger.Error("Error validating required fields:" + err.Error())
		resp.HandleError(ctx, http.StatusUnprocessableEntity, "Unprocessable Entity", nil)
		return
	}

	markerModel := model.Marker{
		Description: markerRequest.Description,
		Latitude:    markerRequest.Latitude,
		Longitude:   markerRequest.Longitude,
		Icon:        markerRequest.Icon,
	}

	err := m.markerService.Create(&markerModel)
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	resp.HandleSuccess(ctx, markerModel)
}

func (m *markerHandler) Delete(ctx *gin.Context) {
	panic("unimplemented")
}

func (*markerHandler) GetAll(ctx *gin.Context) {
	panic("unimplemented")
}

func (*markerHandler) GetByID(ctx *gin.Context) {
	panic("unimplemented")
}

func (*markerHandler) Update(ctx *gin.Context) {
	panic("unimplemented")
}

func NewMarkerHandler(handler *Handler, markerService service.MarkerService) MarkerHandler {
	return &markerHandler{
		Handler:       handler,
		markerService: markerService,
	}
}
