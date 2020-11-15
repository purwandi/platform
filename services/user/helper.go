package user

import (
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/migrator"
	"github.com/purwandi/platform/services/user/migrations"
)

// Migrator service
func Migrator(m *migrator.Migrator) {
	m.Register(20201114001, migrations.MigrateCreateUser, migrations.RollbackCreateUser)
}

// Faker ...
func Faker(repo rel.Repository) {

}
