package resolver

import (
	"context"
	"errors"

	"github.com/purwandi/platform/gateway/types"
	"github.com/purwandi/platform/services/module"
	"github.com/spf13/cast"
)

// ModuleTypes is for get avaiable module types
func (r *Resolver) ModuleTypes(ctx context.Context) ([]types.ModuleType, error) {
	mt := []types.ModuleType{}

	// Process
	res, err := r.ModuleService.FindAllTypes(ctx)
	if err != nil {
		return mt, err
	}

	// Parse into types
	for _, item := range res {
		mt = append(mt, types.ModuleType{Type: item})
	}

	return mt, nil
}

// CreateModuleInput for create module input
type CreateModuleInput struct {
	OrderID     *int32
	ProjectID   int32
	TypeID      string
	Location    *string
	Description *string
}

// CreateModule for creating module
func (r *Resolver) CreateModule(ctx context.Context, args struct{ Input CreateModuleInput }) (types.Module, error) {
	u, err := r.AuthService.User(ctx)
	if err != nil {
		return types.Module{}, err
	}

	// Validate type module
	t, err := r.ModuleService.FindTypeByCode(ctx, args.Input.TypeID)
	if err != nil {
		return types.Module{}, err
	}

	// Check product is exists
	p, err := r.ProductService.FindByID(ctx, cast.ToInt(args.Input.ProjectID))
	if err != nil {
		return types.Module{}, err
	}

	// Validate if project created by current id
	if p.UserID != u.ID {
		return types.Module{}, errors.New("Unathorized for creating module")
	}

	// Given
	inpt := module.CreateModuleInput{
		ProjectID:   p.ID,
		TypeID:      t.Code,
		Location:    args.Input.Location,
		Description: args.Input.Description,
	}

	// Process
	m, err := r.ModuleService.CreateModule(ctx, inpt)
	if err != nil {
		return types.Module{}, err
	}

	return types.Module{Module: m}, nil
}
