package migrations

import (
	"github.com/go-rel/rel"
)

// MigrateCreateProjectRole ...
func MigrateCreateProjectRole(schema *rel.Schema) {
	schema.CreateTable("project_roles", func(t *rel.Table) {
		t.ID("id")
		t.String("name", rel.Limit(64))
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})

	schema.CreateTable("user_project_roles", func(t *rel.Table) {
		t.Int("user_id", rel.Unsigned(true))
		t.Int("project_role_id", rel.Unsigned(true))

		t.ForeignKey("user_id", "users", "id")
		t.ForeignKey("project_role_id", "project_roles", "id")
	})
}

// RollbackCreateProjectRole ...
func RollbackCreateProjectRole(schema *rel.Schema) {
	schema.DropTable("user_project_roles")
	schema.DropTable("project_role")
}
