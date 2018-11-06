package main

import (
	"os"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-varnish/src/args"
	"github.com/newrelic/nri-varnish/src/metrics"
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

	entity := i.LocalEntity()

	// Collect inventory from files
	if args.HasInventory() {
		collectInventory(entity, &args)
	}

	if args.HasMetrics() {
		if err := metrics.CollectMetrics(entity, i); err != nil {
			log.Error("Error collecting metrics: %s", err.Error())
			os.Exit(1)
		}
	}

	if err = i.Publish(); err != nil {
		log.Error(err.Error())
	}
}
