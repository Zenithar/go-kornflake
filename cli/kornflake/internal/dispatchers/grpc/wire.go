//+build wireinject

package grpc

import (
	"context"
	"crypto/tls"

	"go.zenithar.org/kornflake/cli/kornflake/internal/config"
	"go.zenithar.org/kornflake/cli/kornflake/internal/core"
	v1 "go.zenithar.org/kornflake/internal/services/v1"
	bigflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/bigflake/v1"
	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"

	"github.com/google/wire"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/tlsconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func grpcServer(ctx context.Context, cfg *config.Configuration, snowflakes v1.SnowflakeGenerator, bigflakes v1.BigflakeGenerator) (*grpc.Server, error) {
	// gRPC middlewares
	sopts := []grpc.ServerOption{}

	// Replace gRPC logger
	grpc_zap.ReplaceGrpcLogger(zap.L())

	// gRPC middlewares
	sopts = append(sopts, grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zap.L()),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(zap.L()),
			),
		),
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)

	// Enable TLS if requested
	if cfg.Server.GRPC.UseTLS {
		// Client authentication enabled but not required
		clientAuth := tls.VerifyClientCertIfGiven
		if cfg.Server.GRPC.TLS.ClientAuthenticationRequired {
			clientAuth = tls.RequireAndVerifyClientCert
		}

		// Generate TLS configuration
		tlsConfig, err := tlsconfig.Server(tlsconfig.Options{
			KeyFile:    cfg.Server.GRPC.TLS.PrivateKeyPath,
			CertFile:   cfg.Server.GRPC.TLS.CertificatePath,
			CAFile:     cfg.Server.GRPC.TLS.CACertificatePath,
			ClientAuth: clientAuth,
		})
		if err != nil {
			log.For(ctx).Error("Unable to build TLS configuration from settings", zap.Error(err))
			return nil, err
		}

		// Create the TLS credentials
		sopts = append(sopts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	} else {
		log.For(ctx).Info("No transport authentication enabled")
	}

	// Initialize the server
	server := grpc.NewServer(sopts...)

	// Health
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)

	// Register services
	snowflakev1.RegisterSnowflakeAPIServer(server, snowflakes)
	bigflakev1.RegisterBigflakeAPIServer(server, bigflakes)

	// Reflection
	reflection.Register(server)

	// Register stat views
	err := view.Register(
		// HTTP
		ochttp.ServerRequestCountView,
		ochttp.ServerRequestBytesView,
		ochttp.ServerResponseBytesView,
		ochttp.ServerLatencyView,
		ochttp.ServerRequestCountByMethod,
		ochttp.ServerResponseCountByStatusCode,
	)
	if err != nil {
		log.For(ctx).Fatal("Unable to register HTTP stat views", zap.Error(err))
	}

	err = view.Register(ocgrpc.DefaultServerViews...)
	if err != nil {
		log.For(ctx).Fatal("Unable to register gRPC stat views", zap.Error(err))
	}

	// Return no error
	return server, nil
}

// -----------------------------------------------------------------------------

func setup(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {
	wire.Build(
		core.V1ServiceSet,
		grpcServer,
	)
	return &grpc.Server{}, nil
}
