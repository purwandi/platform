package cmd

import (
	"github.com/go-rel/rel/migrator"
	"github.com/purwandi/platform/services/user"
	"github.com/spf13/cobra"
)

func migrate(operation string) {
	m := migrator.New(relw)

	// register migrator service
	user.Migrator(&m)

	switch operation {
	default:
		logger.Error("Unrecognized db migration command")
	case "migrate":
		m.Migrate(ctx)
	case "rollback":
		m.Rollback(ctx)
	}

	// close
	for i := range shutdowns {
		shutdowns[i]()
	}
}

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the last database migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate("migrate")
	},
}

var dbRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback the last database migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate("rollback")
	},
}

var dbSeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with records",
	Run: func(cmd *cobra.Command, args []string) {

		// Call user seeder
		user.Seeder(ctx, relw)

		// close
		for i := range shutdowns {
			shutdowns[i]()
		}
	},
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Run database migration",
}

func init() {
	dbCmd.AddCommand(dbMigrateCmd)
	dbCmd.AddCommand(dbRollbackCmd)
	dbCmd.AddCommand(dbSeedCmd)
}
