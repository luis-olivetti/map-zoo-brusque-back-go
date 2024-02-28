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
	markerHandler handler.MarkerHandler,
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

	router.POST("/user/login", userHandler.Login)

	router.GET("/marker", markerHandler.GetAll)
	router.GET("/marker/:id", markerHandler.GetByID)
	router.POST("/marker", markerHandler.Create)
	router.PUT("/marker/:id", markerHandler.Update)
	router.DELETE("/marker/:id", markerHandler.Delete)

	return router
}
