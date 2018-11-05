package main

import (
	"os"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-varnish/src/args"
)

const (
	integrationName    = "com.newrelic.varnish"
	integrationVersion = "0.1.0"
)

func main() {
	var args args.ArgumentList

	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	// Setup logging with verbose
	log.SetupLogging(args.Verbose)

	// Collect inventory from files
	if args.HasInventory() {

	}

	// entity := i.LocalEntity()

	if err = i.Publish(); err != nil {
		log.Error(err.Error())
	}
}
