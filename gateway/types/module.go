package types

import (
	"github.com/purwandi/platform/services/module"
	"github.com/spf13/cast"
)

// Module entity
type Module struct {
	module.Module
}

// ID ...
func (m Module) ID() int32 {
	return cast.ToInt32(m.Module.ID)
}

// OrderID ...
func (m Module) OrderID() *int32 {
	or := cast.ToInt32(m.Module.OrderID)
	return &or
}

// ProjectID ...
func (m Module) ProjectID() int32 {
	return cast.ToInt32(m.Module.ProjectID)
}
