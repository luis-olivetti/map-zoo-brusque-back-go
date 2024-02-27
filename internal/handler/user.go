package handler

import (
	"github.com/gin-gonic/gin"
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
	h.logger.Info("Login")

	r, _ := h.userService.GenerateJWT("test", "test")

	resp.HandleSuccess(ctx, r)
}
