package grpc

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"

	"go.zenithar.org/kornflake/cli/kornflake/internal/config"
	"go.zenithar.org/pkg/log"
)

type application struct {
	cfg    *config.Configuration
	server *grpc.Server
}

var (
	app  *application
	once sync.Once
)

// -----------------------------------------------------------------------------

// New initialize the application
func New(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {
	var err error

	once.Do(func() {
		// Initialize application
		app = &application{}

		// Apply configuration
		if err := app.ApplyConfiguration(cfg); err != nil {
			log.For(ctx).Fatal("Unable to initialize server settings", zap.Error(err))
		}
	})

	app.server, err = setup(ctx, cfg)
	if err != nil {
		return nil, xerrors.Errorf("grpc: unable to initialize core services : %w", err)
	}

	// Return server
	return app.server, nil
}

// -----------------------------------------------------------------------------

// ApplyConfiguration apply the configuration after checking it
func (s *application) ApplyConfiguration(cfg interface{}) error {
	// Check configuration validity
	if err := s.checkConfiguration(cfg); err != nil {
		return err
	}

	// Apply to current component (type assertion done if check)
	s.cfg, _ = cfg.(*config.Configuration)

	// No error
	return nil
}

// -----------------------------------------------------------------------------

func (s *application) checkConfiguration(cfg interface{}) error {
	// No error
	return nil
}
