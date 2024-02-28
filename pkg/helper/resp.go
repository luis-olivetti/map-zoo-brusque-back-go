package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Message: "success", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Message: message, Data: data}
	ctx.JSON(httpCode, resp)
}
