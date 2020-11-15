package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger    *zap.Logger
	dbw       rel.Repository
	dbr       rel.Repository
	shutdowns []func() error

	ctx = context.Background()
)

func initDatabase(dsn string) rel.Repository {
	adapter, err := postgres.Open(dsn)
	if err != nil {
		logger.Fatal("Unable to connect database")
	}

	// register into shutdown function
	shutdowns = append(shutdowns, adapter.Close)

	// creating new
	repo := rel.New(adapter)
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
	dbw = initDatabase(os.Getenv("DATABASE_URL"))
	dbr = initDatabase(os.Getenv("DATABASE_URL"))

	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(apiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
