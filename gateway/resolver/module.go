package resolver

import (
	"context"

	"github.com/purwandi/platform/services/module"
	"github.com/spf13/cast"
)

func (r *Resolver) CreateModule(ctx context.Context, args struct{ Input CreateModuleInput }) (ModuleResolver, error) {
	return ModuleResolver{}, nil
}

type CreateModuleInput struct {
	OrderID     *int32
	ProjectID   int32
	TypeID      string
	Location    *string
	Description *string
}

type ModuleResolver struct {
	module.Module
}

// ID ...
func (m ModuleResolver) ID() int32 {
	return cast.ToInt32(m.Module.ID)
}

// OrderID ...
func (m ModuleResolver) OrderID() *int32 {
	or := cast.ToInt32(m.Module.OrderID)
	return &or
}

// ProjectID ...
func (m ModuleResolver) ProjectID() int32 {
	return cast.ToInt32(m.Module.ProjectID)
}
