package migrations

import (
	"github.com/go-rel/rel"
)

// MigrateCreateRole ...
func MigrateCreateRole(schema *rel.Schema) {
	schema.CreateTable("roles", func(t *rel.Table) {
		t.ID("id")
		t.String("name", rel.Limit(64))
		t.Text("description")
		t.Text("requirement")
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})
}

// RollbackCreateRole ...
func RollbackCreateRole(schema *rel.Schema) {
	schema.DropTable("roles")
}
