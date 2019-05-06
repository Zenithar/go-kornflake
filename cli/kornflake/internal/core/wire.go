package core

import (
	"github.com/google/wire"

	"go.zenithar.org/kornflake/internal/services/v1/snowflake"
)

// V1ServiceSet is an object provider for wire
var V1ServiceSet = wire.NewSet(
	snowflake.New,
)
