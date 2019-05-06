package v1

import (
	"context"

	bigflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/bigflake/v1"
	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"
)

// SnowflakeGenerator defines identity generator service contract
type SnowflakeGenerator interface {
	Get(ctx context.Context, req *snowflakev1.GetRequest) (*snowflakev1.GetResponse, error)
}

// BigflakeGenerator defines identity generator service contract
type BigflakeGenerator interface {
	Get(ctx context.Context, req *bigflakev1.GetRequest) (*bigflakev1.GetResponse, error)
}
