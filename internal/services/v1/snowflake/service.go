package snowflake

import (
	"context"

	"github.com/sony/sonyflake"
	"golang.org/x/xerrors"

	v1 "go.zenithar.org/kornflake/internal/services/v1"
	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"
)

type service struct {
	generator *sonyflake.Sonyflake
}

// New service implementation using sonyflake
func New(machineID uint16) v1.IdentifierGenerator {
	return &service{
		generator: sonyflake.NewSonyflake(sonyflake.Settings{}),
	}
}

// -----------------------------------------------------------------------------

func (s *service) Get(ctx context.Context, _ *snowflakev1.GetRequest) (*snowflakev1.GetResponse, error) {
	// Generate an id
	id, err := s.generator.NextID()
	if err != nil {
		return nil, xerrors.Errorf("snowflake: unable to generate an identifier : %w", err)
	}

	// Return result
	return &snowflakev1.GetResponse{
		Identifier: id,
	}, nil
}
