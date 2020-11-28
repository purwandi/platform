package resolver

import (
	"context"
	"errors"

	"github.com/purwandi/platform/gateway/types"
	"github.com/purwandi/platform/services/product"
	"github.com/spf13/cast"
)

// CreateProductInput ...
type CreateProductInput struct {
	Code        *string
	Name        string
	Description *string
}

// UpdateProductInput for update product input
type UpdateProductInput struct {
	ID          int32
	Code        *string
	Name        string
	Description *string
}

// CreateProduct ...
func (r *Resolver) CreateProduct(ctx context.Context, args struct{ Input CreateProductInput }) (types.Product, error) {
	// Get current user
	u, err := r.AuthService.User(ctx)
	if err != nil {
		return types.Product{}, err
	}

	// Given
	inpt := product.CreateProductInput{
		UserID:      u.ID,
		Code:        args.Input.Code,
		Name:        args.Input.Name,
		Description: args.Input.Description,
	}

	// Process
	p, err := r.ProductService.Create(ctx, inpt)
	if err != nil {
		return types.Product{}, err
	}

	// Return
	return types.Product{Product: p}, nil
}

// UpdateProduct ...
func (r *Resolver) UpdateProduct(ctx context.Context, args struct{ Input UpdateProductInput }) (types.Product, error) {
	u, err := r.AuthService.User(ctx)
	if err != nil {
		return types.Product{}, err
	}

	// Given
	inpt := product.UpdateProductInput{
		ID:          cast.ToInt(args.Input.ID),
		Code:        args.Input.Code,
		Name:        args.Input.Name,
		Description: args.Input.Description,
	}

	// Process
	p, err := r.ProductService.FindByID(ctx, inpt.ID)
	if err != nil {
		return types.Product{Product: p}, err
	}

	// check product ownership
	if p.UserID != u.ID {
		return types.Product{}, errors.New("Is not allowed")
	}

	// Update
	p, err = r.ProductService.Update(ctx, p, inpt)
	if err != nil {
		return types.Product{}, err
	}

	return types.Product{Product: p}, nil
}
