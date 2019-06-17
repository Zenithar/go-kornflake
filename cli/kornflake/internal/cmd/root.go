package cmd

import (
	iconfig "go.zenithar.org/kornflake/cli/kornflake/internal/config"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/config"
	cmdcfg "go.zenithar.org/pkg/config/cmd"
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
)

// -----------------------------------------------------------------------------

// RootCmd describes root command of the tool
var mainCmd = &cobra.Command{
	Use:   "kornflake",
	Short: "Distributed identifier generator",
}

func init() {
	mainCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (config.toml)")
	mainCmd.AddCommand(versionCmd)
	mainCmd.AddCommand(cmdcfg.NewConfigCommand(conf, "KRNF"))
	mainCmd.AddCommand(serverCmd)
	mainCmd.AddCommand(clientCmd)
}

// -----------------------------------------------------------------------------

// Execute main command
func Execute() error {
	feature.DefaultMutableGate.AddFlag(mainCmd.Flags())
	return mainCmd.Execute()
}

// -----------------------------------------------------------------------------

var (
	cfgFile string
	conf    = &iconfig.Configuration{}
)

// -----------------------------------------------------------------------------

func initConfig() {
	if err := config.Load(conf, "KRNF", cfgFile); err != nil {
		log.Bg().Fatal("Unable load config", zap.Error(err))
	}
}
