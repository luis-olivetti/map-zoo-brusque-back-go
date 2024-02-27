package handler

import "github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/log"

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
