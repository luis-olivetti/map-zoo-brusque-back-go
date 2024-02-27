package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/middleware"
	resp "github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/helper"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Oliv8!",
		})
	})

	return r
}
