package kernel

import (
	"github.com/go-rel/rel"
	"github.com/purwandi/platform/services/module"
	"github.com/purwandi/platform/services/product"
	"github.com/purwandi/platform/services/user"
)

// Service ...
type Service struct {
	AuthService    *user.AuthService
	UserService    *user.Service
	ProductService *product.Service
	ModuleService  *module.Service
}

// InitService ...
func InitService(repo rel.Repository) Service {
	s := Service{}
	s.AuthService = user.NewAuthService()
	s.UserService = user.NewService(repo)
	s.ProductService = product.NewService(repo, s.UserService)
	s.ModuleService = module.NewService(repo)

	return s
}
