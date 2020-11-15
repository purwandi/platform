package migrations

import "github.com/go-rel/rel"

// MigrateCreateUser ...
func MigrateCreateUser(schema *rel.Schema) {
	schema.CreateTable("users", func(t *rel.Table) {
		t.ID("id")
		t.String("email", rel.Limit(64))
		t.String("password")
		t.String("access_token", rel.Limit(128))
		t.String("fullname", rel.Limit(128))
		t.String("callname", rel.Limit(32))
		t.Bool("active", rel.Default(false))
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})
}

// RollbackCreateUser ...
func RollbackCreateUser(schema *rel.Schema) {
	schema.DropTable("users")
}
