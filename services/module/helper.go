package module

import (
	"github.com/go-rel/rel/migrator"
	"github.com/purwandi/platform/services/module/migrations"
)

func Migrator(m *migrator.Migrator) {
	m.Register(20201123001, migrations.MigrateCreateModule, migrations.RollbackCreateModule)
}
