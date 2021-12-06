package cmd

import (
	log "github.com/sirupsen/logrus"
)

// LogLevel is a flag that when parsed correctly will set the log level in
// logrus.
type LogLevel string

func (level *LogLevel) UnmarshalFlag(value string) error {
	switch value {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	}

	stored := LogLevel(value)
	*level = stored
	return nil
}
