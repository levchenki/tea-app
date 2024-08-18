package service

import (
	"fmt"
	"github.com/levchenki/tea-app/internal/entity"
)

type TeaRepository interface {
	Get() ([]entity.Tea, error)
	GetById(id int) (entity.Tea, error)
	GetByCategoryId(id int) ([]entity.Tea, error)
}

type TeaService struct {
	r TeaRepository
}

func NewTeaService(r TeaRepository) *TeaService {
	return &TeaService{r: r}
}

func (s *TeaService) Get() ([]entity.Tea, error) {
	teas, err := s.r.Get()
	if err != nil {
		return []entity.Tea{}, fmt.Errorf("failed to get all teas: %w", err)
	}
	return teas, nil
}

func (s *TeaService) GetByTeaId(id int) (entity.Tea, error) {
	tea, err := s.r.GetById(id)
	if err != nil {
		return entity.Tea{}, fmt.Errorf("failed to get tea by id: %w", err)
	}
	return tea, nil
}

func (s *TeaService) GetByCategoryId(id int) ([]entity.Tea, error) {
	teas, err := s.r.GetByCategoryId(id)
	if err != nil {
		return []entity.Tea{}, fmt.Errorf("failed to get all teas by category id: %w", err)
	}
	return teas, nil
}
