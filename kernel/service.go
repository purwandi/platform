package kernel

import (
	"github.com/go-rel/rel"
	"github.com/purwandi/platform/services/user"
)

// Service ...
type Service struct {
	UserService *user.Service
}

// InitService ...
func InitService(repo rel.Repository) *Service {
	return &Service{
		UserService: user.NewService(repo),
	}
}
