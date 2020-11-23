package kernel

import (
	"github.com/go-rel/rel"
	"github.com/purwandi/platform/services/product"
	"github.com/purwandi/platform/services/user"
)

// Service ...
type Service struct {
	AuthService    *user.AuthService
	UserService    *user.Service
	ProductService *product.Service
}

// InitService ...
func InitService(repo rel.Repository) Service {
	s := Service{}
	s.AuthService = user.NewAuthService()
	s.UserService = user.NewService(repo)
	s.ProductService = product.NewService(repo, s.UserService)

	return s
}
