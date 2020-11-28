package resolver

import (
	"context"

	"github.com/purwandi/platform/gateway/types"
)

// Viewer ...
func (r *Resolver) Viewer(ctx context.Context) (types.User, error) {

	// result := types.User{User: user}
	// fmt.Println(result.User.Username)
	// fmt.Println(result.Username)

	return types.User{}, nil
}

// UpdateUserProjectRole ...
func (r *Resolver) UpdateUserProjectRole(ctx context.Context, args struct{ Input []int32 }) (types.User, error) {
	return types.User{}, nil
}
