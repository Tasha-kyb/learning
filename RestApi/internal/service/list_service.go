package service

import (
	"errors"
	"fmt"

	"RestApi/internal/domain"
)

var ErrValidation = errors.New("VALIDATION_FAILED")

type Repository interface {
	Create(title string) (domain.List, error)
	GetByID(id string) (domain.List, error)
	Update(id string, title string) (domain.List, error)
	Delete(id string) error
	List(limit, offset int) ([]domain.List, int, error)
}

type ListService struct {
	repo Repository
}

func NewListService(repo Repository) *ListService {
	return &ListService {repo : repo}
}

func (l *ListService) Create(title string) (domain.List, error) {
	if err := validateTitle(title); err != nil {
		return domain.List{}, err
	}
	return l.repo.Create(title)
}

func (l *ListService) Get(id string) (domain.List, error) {
	return l.repo.GetByID(id)
}

func (l *ListService) Update(id string, title string) (domain.List, error) {
	if err := validateTitle(title); err != nil {
		return domain.List{}, err
	}
	return l.repo.Update(id, title)
}

func (l *ListService) Delete(id string) error {
	return l.repo.Delete(id)
}

func (l *ListService) List(limit, offset int) ([]domain.List, int, error) {
	return l.repo.List(limit, offset)
}

func validateTitle(title string) error {
	if len(title) == 0 || len(title) > 100 {
		return fmt.Errorf("%w: title must be 1..100 chars", ErrValidation)
	}
	return nil
}