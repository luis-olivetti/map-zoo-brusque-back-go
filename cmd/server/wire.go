//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/handler"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/server"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/service"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/log"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		ServiceSet,
		HandlerSet,
	))
}
