package types

import (
	"github.com/purwandi/platform/services/module"
	"github.com/spf13/cast"
)

// ModuleResolver entity
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
