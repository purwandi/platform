package module

import (
	"context"
	"time"

	"github.com/go-rel/rel"
)

// Service ...
type Service struct {
	repository rel.Repository
}

// NewService ....
func NewService(repo rel.Repository) *Service {
	return &Service{repository: repo}
}

// CreateModule ...
func (s *Service) CreateModule(ctx context.Context, inpt CreateModuleInput) (Module, error) {
	// Validate
	if err := inpt.Validate(); err != nil {
		return Module{}, err
	}

	// Set
	m := Module{
		ProjectID:   inpt.ProjectID,
		TypeID:      inpt.TypeID,
		Status:      DRAFT,
		Location:    inpt.Location,
		Description: inpt.Description,
		CreatedAt:   time.Now(),
	}

	// Persist
	if err := s.repository.Insert(ctx, &m); err != nil {
		return m, err
	}

	// return
	return m, nil
}

// Archived ...
func (s *Service) Archived(ctx context.Context, m Module) (Module, error) {
	changeset := rel.NewChangeset(&m)

	m.Status = ARCHIVED
	m.UpdatedAt = time.Now()

	// Persist
	if err := s.repository.Update(ctx, &m, changeset); err != nil {
		return m, nil
	}

	// return
	return m, nil
}

// Published ...
func (s *Service) Published(ctx context.Context, m Module) (Module, error) {
	changeset := rel.NewChangeset(&m)

	m.Status = PUBLISHED
	m.UpdatedAt = time.Now()

	// Persist
	if err := s.repository.Update(ctx, &m, changeset); err != nil {
		return m, nil
	}

	// return
	return m, nil
}

// Postponed ...
func (s *Service) Postponed(ctx context.Context, m Module) (Module, error) {
	changeset := rel.NewChangeset(&m)

	m.Status = POSTPONED
	m.UpdatedAt = time.Now()

	// Persist
	if err := s.repository.Update(ctx, &m, changeset); err != nil {
		return m, nil
	}

	// return
	return m, nil
}

// Ongoing ...
func (s *Service) Ongoing(ctx context.Context, m Module) (Module, error) {
	changeset := rel.NewChangeset(&m)

	m.Status = ONGOING
	m.UpdatedAt = time.Now()

	// Persist
	if err := s.repository.Update(ctx, &m, changeset); err != nil {
		return m, nil
	}

	// return
	return m, nil
}

// Closed ...
func (s *Service) Closed(ctx context.Context, m Module) (Module, error) {
	changeset := rel.NewChangeset(&m)

	m.Status = CLOSED
	m.UpdatedAt = time.Now()

	// Persist
	if err := s.repository.Update(ctx, &m, changeset); err != nil {
		return m, nil
	}

	// return
	return m, nil
}
