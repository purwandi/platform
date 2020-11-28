package types

import (
	"github.com/purwandi/platform/services/module"
)

// ModuleType entity
type ModuleType struct {
	module.Type
}

// Code for cast code
func (m ModuleType) Code() string {
	return string(m.Type.Code)
}
