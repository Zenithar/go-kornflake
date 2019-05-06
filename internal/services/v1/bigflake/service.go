package bigflake

import (
	"context"

	"github.com/mattheath/kala/bigflake"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"

	v1 "go.zenithar.org/kornflake/internal/services/v1"
	bigflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/bigflake/v1"
)

type service struct {
	generator *bigflake.Bigflake
}

// New service implementation using sonyflake
func New() v1.BigflakeGenerator {
	bf, err := bigflake.New(0)
	if err != nil {
		panic(err)
	}

	return &service{
		generator: bf,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Get(ctx context.Context, _ *bigflakev1.GetRequest) (*bigflakev1.GetResponse, error) {
	_, span := trace.StartSpan(ctx, "kornflake.v1.bigflake.Get")
	defer span.End()

	// Generate an id
	id, err := s.generator.Mint()
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, xerrors.Errorf("bigflake: unable to generate an identifier : %w", err)
	}

	// Return result
	return &bigflakev1.GetResponse{
		Identifier: id.String(),
	}, nil
}
