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

	LogLevel LogLevel `long:"log-level" default:"INFO" description:"Log level of output" choice:"TRACE" choice:"DEBUG" choice:"INFO" choice:"WARN" choice:"ERROR" choice:"FATAL" choice:"PANIC"`
}

func (e *Execute) Execute(args []string) error {
	log.WithField("Path", e.Positional.ConfigPath).Info("provided config path")
	return nil
}
