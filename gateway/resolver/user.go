package resolver

import (
	"context"
)

// Viewer ...
func (r *Resolver) Viewer(ctx context.Context) (UserResolver, error) {

	// result := UserResolver{User: user}
	// fmt.Println(result.User.Username)
	// fmt.Println(result.Username)

	return UserResolver{}, nil
}

// UserResolver struct ...
type UserResolver struct {
}

// ID ...
func (u UserResolver) ID() string {
	return ""
}
