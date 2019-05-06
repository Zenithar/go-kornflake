package cmd

import (
	"context"

	"github.com/cloudflare/tableflip"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.zenithar.org/kornflake/cli/kornflake/internal/dispatchers/http"
	"go.zenithar.org/kornflake/internal/version"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform"
)

// -----------------------------------------------------------------------------

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Starts the kornflake HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Initialize config
		initConfig()

		// Starting banner
		log.For(ctx).Info("Starting kornflake HTTP server ...")

		// Start goroutine group
		err := platform.Run(ctx, &platform.Application{
			Debug:           conf.Debug.Enable,
			Name:            "kornflake-http",
			Version:         version.Version,
			Revision:        version.Revision,
			Instrumentation: conf.Instrumentation,
			Builder: func(upg *tableflip.Upgrader, group *run.Group) {
				ln, err := upg.Fds.Listen(conf.Server.HTTP.Network, conf.Server.HTTP.Listen)
				if err != nil {
					log.For(ctx).Fatal("Unable to start HTTP server", zap.Error(err))
				}

				server, err := http.New(ctx, conf)
				if err != nil {
					log.For(ctx).Fatal("Unable to start HTTP server", zap.Error(err))
				}

				group.Add(
					func() error {
						log.For(ctx).Info("Starting HTTP server", zap.Stringer("address", ln.Addr()))
						return server.Serve(ln)
					},
					func(e error) {
						log.For(ctx).Info("Shutting HTTP server down")
						log.SafeClose(server, "Unable to close HTTP server")
					},
				)
			},
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}
