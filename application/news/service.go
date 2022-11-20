package news

import (
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	Service interface {
		Create(data *dto.News) error
		GetAll(data *dto.NewsSlice) error
		GetById(id string, data *dto.News) error
		Delete(id string) error
		Save(data *dto.News) error
	}
	service struct {
		shared shared.Holder
	}
)

func (s *service) Create(data *dto.News) error {
	err := s.shared.DB.Create(data).Error
	return err
}

func (s *service) GetAll(data *dto.NewsSlice) error {
	err := s.shared.DB.Find(data).Error
	return err
}

func (s *service) GetById(id string, data *dto.News) error {
	err := s.shared.DB.Where("id = ?", id).First(data).Error
	return err
}

func (s *service) Delete(id string) error {
	err := s.shared.DB.Delete(&dto.News{}, id).Error
	return err
}

func (s *service) Save(data *dto.News) error {
	err := s.shared.DB.Save(data).Error
	return err
}

func NewNewsService(shared shared.Holder) Service {
	return &service{
		shared: shared,
	}
}
