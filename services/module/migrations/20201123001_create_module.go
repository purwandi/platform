package migrations

import "github.com/go-rel/rel"

// MigrateCreateModule ...
func MigrateCreateModule(schema *rel.Schema) {
	schema.CreateTable("modules", func(t *rel.Table) {
		t.ID("id")
		t.Int("order_id")
		t.Int("project_id", rel.Unsigned(true))
		t.String("type_id", rel.Limit(2))
		t.String("status", rel.Limit(16), rel.Default("DRAFT"))
		t.String("location")
		t.Float("success_rate", rel.Precision(5), rel.Scale(2))
		t.Text("description")
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})

	schema.CreateIndex("modules", "modules_project_id", []string{"project_id", "type_id"})
}

// RollbackCreateModule ...
func RollbackCreateModule(schema *rel.Schema) {
	schema.DropTable("modules")
}
