// Package args contains the argument list, defined as a struct, along with a method that validates passed-in args
package args

import (
	"errors"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
)

// ArgumentList struct that holds all Varnish arguments
type ArgumentList struct {
	sdkArgs.DefaultArgumentList
	ParamsConfigFile string `default:"" help:"The location of the varnish.params configuration file. If omitted will search in typical install locations"`
	InstanceName     string `default:"" help:"User defined name to identify data from this instance in New Relic"`
	ShowVersion      bool   `default:"false" help:"Print build information and exit"`
}

// Validate validates argument parameters
func (al ArgumentList) Validate() error {
	if al.InstanceName == "" {
		return errors.New("instance_name must not be blank")
	}

	return nil
}
