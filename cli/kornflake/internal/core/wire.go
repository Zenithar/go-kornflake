package core

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"go.zenithar.org/kornflake/cli/kornflake/internal/config"
	v1 "go.zenithar.org/kornflake/internal/services/v1"
	"go.zenithar.org/kornflake/internal/services/v1/bigflake"
	"go.zenithar.org/kornflake/internal/services/v1/snowflake"
	"go.zenithar.org/pkg/log"
)

// V1ServiceSet is an object provider for wire
var V1ServiceSet = wire.NewSet(
	Snowflake,
	Bigflake,
)

// Snowflake build the generator using configuration file
func Snowflake(cfg *config.Configuration) v1.SnowflakeGenerator {
	const maxUint32 = ^uint32(0)
	if cfg.WorkerID > uint64(maxUint32) {
		log.Bg().Warn("Unable to initialize snowflake generator with this WorkerID", zap.Uint64("worker_id", cfg.WorkerID))
	}
	return snowflake.New(uint32(cfg.WorkerID) % maxUint32)
}

// Bigflake build the generator using configuration file
func Bigflake(cfg *config.Configuration) v1.BigflakeGenerator {
	return bigflake.New(cfg.WorkerID)
}
