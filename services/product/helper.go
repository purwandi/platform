package product

import (
	"github.com/go-rel/rel/migrator"
	"github.com/purwandi/platform/services/product/migrations"
)

func Migrator(m *migrator.Migrator) {
	m.Register(20201122001, migrations.MigrateCreateProduct, migrations.RollbackCreateProduct)
}
