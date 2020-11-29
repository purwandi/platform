package types

import (
	"github.com/purwandi/platform/services/module"
	"github.com/spf13/cast"
)

// Role ...
type Role struct {
	module.Role
}

// ID ...
func (r Role) ID() int32 {
	return cast.ToInt32(r.Role.ID)
}
