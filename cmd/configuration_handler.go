package cmd

import "github.com/goodmustache/appli.ed/configuration"

//counterfeiter:generate . ConfigurationHandler

type ConfigurationHandler interface {
	ReadConfig(path string) (configuration.WatchConfiguration, error)
}
