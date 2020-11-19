package resolver

import (
	"context"

	"github.com/purwandi/platform/services/user"
)

// Viewer ...
func (r *Resolver) Viewer(ctx context.Context) (UserResolver, error) {

	// result := UserResolver{User: user}
	// fmt.Println(result.User.Username)
	// fmt.Println(result.Username)

	return UserResolver{}, nil
}

// UpdateUserProjectRole ...
func (r *Resolver) UpdateUserProjectRole(ctx context.Context, args struct{ Input []int32 }) (UserResolver, error) {
	return UserResolver{}, nil
}

// UserResolver struct ...
type UserResolver struct {
	user.User
}

// ID ...
func (u UserResolver) ID() string {
	return ""
}
