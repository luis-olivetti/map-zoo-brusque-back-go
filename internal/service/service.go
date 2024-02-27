package service

import "github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/log"

type Service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}
