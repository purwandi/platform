package product

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/migrator"
	"github.com/purwandi/platform/services/product/migrations"
)

func Migrator(m *migrator.Migrator) {
	m.Register(20201122001, migrations.MigrateCreateProduct, migrations.RollbackCreateProduct)
}

func Seeder(ctx context.Context, repo rel.Repository) {
	repo.InsertAll(ctx, &[]Product{
		{ID: 1, UserID: 1, Name: "Learning Management System"},
		{ID: 2, UserID: 1, Name: "Ku Laku-Laku"},
	})
}
