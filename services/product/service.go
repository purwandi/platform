package product

import (
	"context"
	"errors"

	"github.com/go-rel/rel"
)

type userService interface {
	UserExists(context.Context, int) bool
}

// Service ...
type Service struct {
	repository  rel.Repository
	userService userService
}

// NewService ...
func NewService(repo rel.Repository, us userService) *Service {
	return &Service{
		repository:  repo,
		userService: us,
	}
}

// CreateProduct ...
func (s *Service) CreateProduct(ctx context.Context, inpt CreateProductInput) (Product, error) {
	// Process
	p, err := CreateProduct(inpt)
	if err != nil {
		return Product{}, err
	}

	// Check if user is exists
	if ok := s.userService.UserExists(ctx, inpt.UserID); ok == false {
		return Product{}, errors.New("User is not exists")
	}

	// Persist
	if err := s.repository.Insert(ctx, p); err != nil {
		return Product{}, err
	}

	// Return
	return *p, nil
}
