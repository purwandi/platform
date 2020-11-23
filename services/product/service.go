package product

import (
	"context"
	"errors"
	"time"

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

// Create ...
func (s *Service) Create(ctx context.Context, inpt CreateProductInput) (Product, error) {
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

// FindByID ...
func (s *Service) FindByID(ctx context.Context, id int) (Product, error) {
	p := Product{}
	err := s.repository.Find(ctx, &p, rel.Eq("id", id))
	if (err != nil) || (p.ID == 0) {
		return p, errors.New("Product not found")
	}

	return p, nil
}

// Update ...
func (s *Service) Update(ctx context.Context, p Product, inpt UpdateProductInput) (Product, error) {
	// Validate ...
	err := inpt.Validate()
	if err != nil {
		return p, err
	}

	// Update
	changeset := rel.NewChangeset(&p)
	p.Code = inpt.Code
	p.Name = inpt.Name
	p.Description = inpt.Description
	p.UpdatedAt = time.Now()

	// Persist
	if err = s.repository.Update(ctx, &p, changeset); err != nil {
		return p, err
	}

	// Return
	return p, nil
}
