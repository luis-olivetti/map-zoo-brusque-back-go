package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			ctx.Header("Access-Control-Allow-Methods", ctx.GetHeader("Access-Control-Request-Method"))
			ctx.Header("Access-Control-Allow-Headers", ctx.GetHeader("Access-Control-Request-Headers"))
			ctx.Header("Access-Control-Max-Age", "7200")
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
