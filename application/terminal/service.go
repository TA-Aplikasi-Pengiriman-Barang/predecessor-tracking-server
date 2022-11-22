package terminal

import (
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	Service interface {
		GetById(id string, data *dto.Terminal) error
		GetAllByRoute(route dto.Route, data *[]dto.Terminal) error
		GetAllTerminal(data *[]dto.Terminal) error
	}
	service struct {
		shared shared.Holder
	}
)

func (s *service) GetById(id string, data *dto.Terminal) error {
	err := s.shared.DB.Where("id = ?", id).First(data).Error
	return err
}

func (s *service) GetAllByRoute(route dto.Route, data *[]dto.Terminal) error {
	err := s.shared.DB.Where("route = ?", route).Find(data).Error
	return err
}

func (s *service) GetAllTerminal(data *[]dto.Terminal) error {
	err := s.shared.DB.Find(data).Error
	return err
}

func NewTerminalService(shared shared.Holder) Service {
	return &service{
		shared: shared,
	}
}
