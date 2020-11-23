package resolver

import (
	"context"

	"github.com/purwandi/platform/services/product"
	"github.com/spf13/cast"
)

// CreateProductInput ...
type CreateProductInput struct {
	Code        *string
	Name        string
	Description *string
}

// CreateProduct ...
func (r *Resolver) CreateProduct(ctx context.Context, args struct{ Input CreateProductInput }) (ProductResolver, error) {
	u, err := r.AuthService.User(ctx)
	if err != nil {
		return ProductResolver{}, err
	}

	// Given
	inpt := product.CreateProductInput{
		UserID:      u.ID,
		Code:        args.Input.Code,
		Name:        args.Input.Name,
		Description: args.Input.Description,
	}

	// Process
	p, err := r.ProductService.CreateProduct(ctx, inpt)
	if err != nil {
		return ProductResolver{}, err
	}

	// Return
	return ProductResolver{Product: p}, nil
}

// ProductResolver ...
type ProductResolver struct {
	product.Product
}

// ID ...
func (p ProductResolver) ID() int32 {
	return cast.ToInt32(p.Product.ID)
}
