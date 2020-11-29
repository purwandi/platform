package module

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/migrator"
	"github.com/purwandi/platform/services/module/migrations"
)

func Migrator(m *migrator.Migrator) {
	m.Register(20201123001, migrations.MigrateCreateModule, migrations.RollbackCreateModule)
	m.Register(20201129001, migrations.MigrateCreateRole, migrations.RollbackCreateRole)
	m.Register(20201129002, migrations.MigrateCreateModuleRole, migrations.RollbackCreateModuleRole)
}

func Seeder(ctx context.Context, repo rel.Repository) {
	// roles
	repo.InsertAll(ctx, &[]Role{
		{ID: 1, Name: "Lead Design Researcher"},
		{ID: 2, Name: "Assistant Researcher"},
		{ID: 3, Name: "Lead Facilitator"},
		{ID: 4, Name: "Co-Facilitator"},
		{ID: 5, Name: "Product Owner"},
		{ID: 6, Name: "Scrum Master"},
		{ID: 7, Name: "Software Tester"},
		{ID: 8, Name: "Lead Technical Analyst"},
		{ID: 9, Name: "Co-Technical Analyst"},
		{ID: 10, Name: "Copy Writer"},
		{ID: 11, Name: "UI Designer"},
		{ID: 12, Name: "UX Designer"},
		{ID: 13, Name: "Backend Programmer"},
		{ID: 14, Name: "Frontend Programmer"},
		{ID: 15, Name: "Devops"},
		{ID: 16, Name: "Photographer"},
	})
}
