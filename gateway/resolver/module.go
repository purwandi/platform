package resolver

import (
	"context"

	"github.com/purwandi/platform/gateway/types"
)

// ModuleTypes is for get avaiable module types
func (r *Resolver) ModuleTypes(ctx context.Context) ([]types.ModuleTypeResolver, error) {
	mt := []types.ModuleTypeResolver{}

	// Process
	res, err := r.ModuleService.FindAllTypes(ctx)
	if err != nil {
		return mt, err
	}

	// Parse into types
	for _, item := range res {
		mt = append(mt, types.ModuleTypeResolver{Type: item})
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
func (r *Resolver) CreateModule(ctx context.Context, args struct{ Input CreateModuleInput }) (types.ModuleResolver, error) {
	return types.ModuleResolver{}, nil
}
