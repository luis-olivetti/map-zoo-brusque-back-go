package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/request"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/service"
	resp "github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/helper"
)

type UserHandler interface {
	Login(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *userHandler) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		h.logger.Error("Invalid request" + err.Error())
		resp.HandleError(ctx, http.StatusBadRequest, "Bad Request", nil)
		return
	}

	authorized, err := h.userService.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil || !authorized {
		h.logger.Error("Unauthorized")
		resp.HandleError(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	response, _ := h.userService.GenerateJWT(loginRequest.Username)

	resp.HandleSuccess(ctx, response)
}
