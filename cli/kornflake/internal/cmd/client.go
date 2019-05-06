package cmd

import (
	"github.com/spf13/cobra"

	bigflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/bigflake/v1"
	snowflakev1 "go.zenithar.org/kornflake/pkg/gen/go/identifier/snowflake/v1"
)

// -----------------------------------------------------------------------------

var clientCmd = &cobra.Command{
	Use:     "client",
	Aliases: []string{"c", "cli"},
	Short:   "Query the gRPC server",
}

func init() {
	clientCmd.AddCommand(
		snowflakev1.SnowflakeAPIClientCommand,
		bigflakev1.BigflakeAPIClientCommand,
	)
}
