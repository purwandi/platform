package resolver

import (
	"context"

	"github.com/purwandi/platform/gateway/types"
)

// Roles ...
func (r *Resolver) Roles(ctx context.Context) ([]types.Role, error) {
	tr := []types.Role{}

	// Process
	res, err := r.ModuleService.FindAllRoles(ctx)
	if err != nil {
		return tr, err
	}

	// Parse
	for _, item := range res {
		tr = append(tr, types.Role{Role: item})
	}

	return tr, nil
}
