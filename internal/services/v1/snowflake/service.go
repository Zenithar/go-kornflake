package snowflake

import (
	"context"

	"github.com/mattheath/kala/snowflake"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"

	v1 "go.zenithar.org/kornflake/internal/services/v1"
	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"
)

type service struct {
	generator *snowflake.Snowflake
}

// New service implementation using sonyflake
func New(workerID uint32) v1.SnowflakeGenerator {
	sf, err := snowflake.New(workerID)
	if err != nil {
		panic(err)
	}
	return &service{
		generator: sf,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Get(ctx context.Context, _ *snowflakev1.GetRequest) (*snowflakev1.GetResponse, error) {
	_, span := trace.StartSpan(ctx, "kornflake.v1.snowflake.Get")
	defer span.End()

	// Generate an id
	id, err := s.generator.Mint()
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, xerrors.Errorf("snowflake: unable to generate an identifier : %w", err)
	}

	// Return result
	return &snowflakev1.GetResponse{
		Identifier: id,
	}, nil
}
