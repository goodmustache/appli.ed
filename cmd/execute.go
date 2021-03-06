package cmd

import (
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

type PositionalArgs struct {
	ConfigPath flags.Filename `positional-arg-name:"PATH_TO_CONFIG" required:"true" description:"Path to the appli.ed configuration file."`
}

type Execute struct {
	Positional PositionalArgs `positional-args:"yes"`

	LogLevel LogLevel `long:"log-level" env:"LOG_LEVEL" default:"INFO" description:"Log level of output" choice:"TRACE" choice:"DEBUG" choice:"INFO" choice:"WARN" choice:"ERROR" choice:"FATAL" choice:"PANIC"`

	ConfigurationHandler ConfigurationHandler
}

func (command *Execute) Execute(args []string) error {
	log.WithField("Path", command.Positional.ConfigPath).Info("provided config path")
	return nil
}
