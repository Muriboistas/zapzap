package user

import (
	"time"

	"github.com/muriboistas/zapzap/entity"
)

// Service interface
type Service struct {
	repo Repository
}

// NewService create a new usecase service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// READERS

// GetUser by urn
func (s *Service) GetUser(urn string) (*entity.User, error) {
	return s.repo.Get(urn)
}

// WRITERS

// CreateUser ...
func (s *Service) CreateUser(name, urn, language string) (entity.ID, error) {
	e, err := entity.NewUser(name, urn, language)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

// UpdateUser ...
func (s *Service) UpdateUser(e *entity.User) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// DeleteUser ...
func (s *Service) DeleteUser(urn string) error {
	u, err := s.GetUser(urn)
	if u == nil {
		return ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(u.ID)
}
