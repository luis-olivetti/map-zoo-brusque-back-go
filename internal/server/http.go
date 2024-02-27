package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/handler"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/middleware"
	resp "github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/helper"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(
		middleware.CORSMiddleware(),
	)
	router.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Oliv8!",
		})
	})

	router.GET("/user/login", userHandler.Login)

	return router
}
