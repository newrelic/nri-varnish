// Package args contains the argument list, defined as a struct, along with a method that validates passed-in args
package args

import (
	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
)

// ArgumentList struct that holds all Varnish arguments
type ArgumentList struct {
	sdkArgs.DefaultArgumentList
	ParamsConfigFile string `default:"" help:"The location of the varnish.params configuration file. If omitted will search in typical install locations"`
}
