package migrations

import (
	"github.com/go-rel/rel"
)

// MigrateCreateModuleRole ...
func MigrateCreateModuleRole(schema *rel.Schema) {
	schema.CreateTable("modules_roles", func(t *rel.Table) {
		t.Int("id")
		t.Int("role_id", rel.Unsigned(true))
		t.Int("quantity", rel.Unsigned(true))
		t.Text("nice_to_have")
		t.Bool("is_fixed_quantity", rel.Default(false))
		t.Bool("is_lead", rel.Default(false))
		t.Int("duration", rel.Unsigned(true))
		t.Int("total_hired", rel.Unsigned(true))
		t.Int("total_confirmed", rel.Unsigned(true))
		t.DateTime("created_at")
		t.DateTime("updated_at")

		t.ForeignKey("role_id", "roles", "id")
	})
}

// RollbackCreateModuleRole ...
func RollbackCreateModuleRole(schema *rel.Schema) {
	schema.DropTable("modules_roles")
}
