package migrations

import (
	"github.com/go-rel/rel"
)

func MigrateCreateProduct(schema *rel.Schema) {
	schema.CreateTable("products", func(t *rel.Table) {
		t.ID("id")
		t.Int("user_id", rel.Unsigned(true))
		t.String("code", rel.Limit(64))
		t.String("name", rel.Limit(128))
		t.Text("description")
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})
}

func RollbackCreateProduct(schema *rel.Schema) {
	schema.DropTable("products")
}
