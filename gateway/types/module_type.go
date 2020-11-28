package types

import (
	"github.com/purwandi/platform/services/module"
	"github.com/spf13/cast"
)

// ModuleTypeResolver entity
type ModuleTypeResolver struct {
	module.Type
}

// ID for map id
func (m ModuleTypeResolver) ID() int32 {
	return cast.ToInt32(m.Type.ID)
}

// Code for cast code
func (m ModuleTypeResolver) Code() string {
	return string(m.Type.Code)
}
