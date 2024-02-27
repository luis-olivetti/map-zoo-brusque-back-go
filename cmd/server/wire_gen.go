// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*gin.Engine, func(), error) {
	handlerHandler := handler.NewHandler(logger)
	serviceService := service.NewService(logger)
	userService := service.NewUserService(serviceService, viperViper)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	engine := server.NewServerHTTP(logger, userHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var ServerSet = wire.NewSet(server.NewServerHTTP)

var ServiceSet = wire.NewSet(service.NewService, service.NewUserService)

var HandlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)
