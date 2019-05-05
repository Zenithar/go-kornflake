package v1

import (
	"context"

	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"
)

// IdentifierGenerator defines identity generator service contract
type IdentifierGenerator interface {
	Get(ctx context.Context, req *snowflakev1.GetRequest) (*snowflakev1.GetResponse, error)
}
