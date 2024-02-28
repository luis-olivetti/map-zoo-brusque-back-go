package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func JWTMiddleware(conf *viper.Viper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if openPath(ctx) {
			ctx.Next()
			return
		}

		header := ctx.GetHeader("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(401)
			return
		}

		tokenString := strings.Split(header, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.GetString("jwt.secret")), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Next()
	}
}

func openPath(ctx *gin.Context) bool {
	pathLogin := strings.Contains(ctx.FullPath(), "user/login")
	pathGetAllMarkers := ctx.Request.Method == "GET" && ctx.FullPath() == "/marker"

	return pathLogin || pathGetAllMarkers
}
