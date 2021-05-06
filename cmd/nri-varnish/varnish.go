//go:generate goversioninfo
package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-varnish/internal/args"
	"github.com/newrelic/nri-varnish/internal/metrics"
)

const (
	integrationName = "com.newrelic.varnish"
)

var (
	integrationVersion = "0.0.0"
	gitCommit          = ""
	buildDate          = ""
)

func main() {
	var args args.ArgumentList

	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	if args.ShowVersion {
		fmt.Printf(
			"New Relic %s integration Version: %s, Platform: %s, GoVersion: %s, GitCommit: %s, BuildDate: %s\n",
			strings.Title(strings.Replace(integrationName, "com.newrelic.", "", 1)),
			integrationVersion,
			fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			runtime.Version(),
			gitCommit,
			buildDate)
		os.Exit(0)
	}

	// Setup logging with verbose
	log.SetupLogging(args.Verbose)

	// validate arguments
	if err := args.Validate(); err != nil {
		log.Error("Error input validation failure: %s", err.Error())
		os.Exit(1)
	}

	systemEntity, err := i.EntityReportedBy("localhost", args.InstanceName, "va-instance")
	if err != nil {
		log.Error("Error creating system entity: %s", err.Error())
		os.Exit(1)
	}

	// Collect inventory from files
	if args.HasInventory() {
		collectInventory(systemEntity, &args)
	}

	if args.HasMetrics() {
		if err := metrics.CollectMetrics(systemEntity, i); err != nil {
			log.Error("Error collecting metrics: %s", err.Error())
			os.Exit(1)
		}
	}

	if err = i.Publish(); err != nil {
		log.Error(err.Error())
	}
}
