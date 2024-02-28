package service

import (
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/model"
)

type MarkerService interface {
	GetAll() ([]*model.Marker, error)
	GetByID(id int) (*model.Marker, error)
	Create(marker *model.Marker) error
	Update(marker *model.Marker) error
	Delete(id int) error
}

type markerService struct {
	*Service
}

func (m *markerService) Create(marker *model.Marker) error {
	//TODO: implement

	m.logger.Logger.Info("Creating marker")
	return nil
}

func (*markerService) Delete(id int) error {
	panic("unimplemented")
}

func (*markerService) GetAll() ([]*model.Marker, error) {
	panic("unimplemented")
}

func (*markerService) GetByID(id int) (*model.Marker, error) {
	panic("unimplemented")
}

func (*markerService) Update(marker *model.Marker) error {
	panic("unimplemented")
}

func NewMarkerService(service *Service) MarkerService {
	return &markerService{
		Service: service,
	}
}
