package types

import (
	"github.com/purwandi/platform/services/module"
)

// ModuleTypeResolver entity
type ModuleTypeResolver struct {
	module.Type
}

// Code for cast code
func (m ModuleTypeResolver) Code() string {
	return string(m.Type.Code)
}
