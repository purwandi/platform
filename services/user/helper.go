package user

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/migrator"
	"github.com/purwandi/platform/pkg/caster"
	"github.com/purwandi/platform/services/user/migrations"
)

// Migrator service
func Migrator(m *migrator.Migrator) {
	m.Register(20201114001, migrations.MigrateCreateUser, migrations.RollbackCreateUser)
	// m.Register(20201118001, migrations.MigrateCreateProjectRole, migrations.RollbackCreateProjectRole)
}

// Seeder ...
func Seeder(ctx context.Context, repo rel.Repository) {
	repo.InsertAll(ctx, &[]User{
		{
			ID:          1,
			Username:    "foobar",
			Email:       "foo@bar.com",
			AccessToken: caster.String("d3f18f8f-ed5e-4350-b849-29424b951f2a"),
			Fullname:    caster.String("Foo bar"),
			Callname:    caster.String("foo"),
		},
		{
			ID:          2,
			Username:    "barfoo",
			Email:       "bar@foo.com",
			AccessToken: caster.String("2d45432a-3dd7-4dea-888c-84e165b3f392"),
			Fullname:    caster.String("Bar Foo"),
			Callname:    caster.String("bar"),
		},
	})
}
