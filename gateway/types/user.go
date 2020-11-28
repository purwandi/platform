package types

import (
	"github.com/purwandi/platform/services/user"
)

// User struct ...
type User struct {
	user.User
}

// ID ...
func (u User) ID() string {
	return ""
}
