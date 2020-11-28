package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	"github.com/go-rel/rel/adapter/sqlite3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger    *zap.Logger
	relw      rel.Repository
	relr      rel.Repository
	shutdowns []func() error

	ctx = context.Background()
)

func initDatabase(conn, dsn string) rel.Repository {
	var repo rel.Repository

	if conn == "postgres" {
		adapter, err := postgres.Open(dsn)
		if err != nil {
			logger.Fatal("Unable to connect database")
		}

		shutdowns = append(shutdowns, adapter.Close)
		repo = rel.New(adapter)
	}

	if conn == "sqlite" {
		adapter, err := sqlite3.Open(dsn)
		if err != nil {
			logger.Fatal("Unable to connect database")
		}

		shutdowns = append(shutdowns, adapter.Close)
		repo = rel.New(adapter)
	}

	repo.Instrumentation(func(ctx context.Context, op, message string) func(err error) {
		if strings.HasPrefix(op, "rel-") {
			return func(error) {}
		}

		t := time.Now()
		return func(err error) {
			d := time.Since(t)
			if err != nil {
				logger.Error(message, zap.Error(err), zap.Duration("duration", d), zap.String("operation", op))
			} else {
				logger.Info(message, zap.Duration("duration", d), zap.String("operation", op))
			}
		}
	})

	return repo
}

var rootCmd = &cobra.Command{
	Use:   "platform",
	Short: "Platform",
}

// Execute command
func Execute() {
	logger, _ = zap.NewProduction(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
		zap.Fields(zap.String("type", "main")),
	)

	// init database connection
	relw = initDatabase(os.Getenv("DB_CONNECTION"), os.Getenv("DB_DSN"))
	relr = initDatabase(os.Getenv("DB_CONNECTION"), os.Getenv("DB_DSN"))

	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(apiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
