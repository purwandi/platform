package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/purwandi/platform/gateway"
	"github.com/purwandi/platform/gateway/resolver"
	"github.com/purwandi/platform/kernel"
	"github.com/purwandi/platform/pkg/server"
	"github.com/purwandi/platform/services/user"
	"github.com/spf13/cobra"
)

type router struct {
	UserServer  *user.Server
	GraphServer *gateway.Server
}

func initAPIServer(srv kernel.Service) router {
	return router{
		UserServer:  user.NewServer(srv.UserService),
		GraphServer: gateway.NewServer(resolver.NewResolver(srv)),
	}
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run api service",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			shutdown = make(chan struct{})
			service  = kernel.InitService(relw)
			api      = initAPIServer(service)
			server   = server.NewServer(ctx)
		)

		// register custom http service component
		server.Register(api.UserServer)
		server.Register(api.GraphServer)

		// register to gracefull context
		shutdowns = append(shutdowns, server.Close)

		go func() {
			quit := make(chan os.Signal)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit

			logger.Info("shutting down server gracefully")

			// close another module
			for i := range shutdowns {
				shutdowns[i]()
			}

			close(shutdown)
		}()

		server.Serve()
		<-shutdown
	},
}
