package resolver

import "github.com/purwandi/platform/kernel"

// Resolver ...
type Resolver struct {
	kernel.Service
}

// NewResolver ...
func NewResolver(service kernel.Service) *Resolver {
	return &Resolver{Service: service}
}
