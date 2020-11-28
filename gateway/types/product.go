package types

import (
	"github.com/purwandi/platform/services/product"
	"github.com/spf13/cast"
)

// Product ...
type Product struct {
	product.Product
}

// ID ...
func (p Product) ID() int32 {
	return cast.ToInt32(p.Product.ID)
}
