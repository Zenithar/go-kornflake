package snowflake

import (
	"context"
	"fmt"
	"time"

	"github.com/sony/sonyflake"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"

	v1 "go.zenithar.org/kornflake/internal/services/v1"
	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"
)

type service struct {
	generator *sonyflake.Sonyflake
}

// New service implementation using sonyflake
func New() v1.IdentifierGenerator {
	return &service{
		generator: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: time.Now().UTC(),
		}),
	}
}

// -----------------------------------------------------------------------------

func (s *service) Get(ctx context.Context, _ *snowflakev1.GetRequest) (*snowflakev1.GetResponse, error) {
	_, span := trace.StartSpan(ctx, "kornflake.v1.snowflake.Get")
	defer span.End()

	// Generate an id
	id, err := s.generator.NextID()
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, xerrors.Errorf("snowflake: unable to generate an identifier : %w", err)
	}

	// Return result
	return &snowflakev1.GetResponse{
		Identifier: fmt.Sprintf("%d", id),
	}, nil
}
