package cmd

import (
	"context"

	"github.com/cloudflare/tableflip"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/kornflake/cli/kornflake/internal/dispatchers/grpc"
	"go.zenithar.org/kornflake/internal/version"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform"
)

// -----------------------------------------------------------------------------

var grpcCmd = &cobra.Command{
	Use:     "grpc",
	Aliases: []string{"g"},
	Short:   "Starts the kornflake gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Initialize config
		initConfig()

		// Start goroutine group
		err := platform.Run(ctx, &platform.Application{
			Debug:           conf.Debug.Enable,
			Name:            "kornflake-grpc",
			Version:         version.Version,
			Revision:        version.Revision,
			Instrumentation: conf.Instrumentation,
			Builder: func(upg *tableflip.Upgrader, group *run.Group) {
				// Starting banner
				log.For(ctx).Info("Starting kornflake gRPC server ...")

				// Allocate listener
				ln, err := upg.Fds.Listen(conf.Server.GRPC.Network, conf.Server.GRPC.Listen)
				if err != nil {
					log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
				}

				// Attach the dispatcher
				server, err := grpc.New(ctx, conf)
				if err != nil {
					log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
				}

				// Add to goroutine group
				group.Add(
					func() error {
						log.For(ctx).Info("GRPC server listening ...", zap.Stringer("address", ln.Addr()))
						return server.Serve(ln)
					},
					func(e error) {
						log.For(ctx).Info("Shutting GRPC server down")
						server.GracefulStop()
					},
				)
			},
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}
