//go:generate goversioninfo
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	args2 "github.com/newrelic/nri-varnish/src/args"
	metrics2 "github.com/newrelic/nri-varnish/src/metrics"

	"github.com/newrelic/infra-integrations-sdk/v3/integration"
	"github.com/newrelic/infra-integrations-sdk/v3/log"
)

// Known locations of varnish.params on specific Linux Distros
const (
	debianUbuntuParamsLoc = "/etc/default/varnish/varnish.params"
	rhelCentosParamsLoc   = "/etc/sysconfig/varnish/varnish.params"
	integrationName       = "com.newrelic.varnish"
)

var (
	integrationVersion = "0.0.0"
	gitCommit          = ""
	buildDate          = ""
)

func main() {
	var args args2.ArgumentList

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
	if err = args.Validate(); err != nil {
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
		if err = metrics2.CollectMetrics(systemEntity, i, args.VarnishName); err != nil {
			log.Error("Error collecting metrics: %s", err.Error())
			os.Exit(1)
		}
	}

	if err = i.Publish(); err != nil {
		log.Error(err.Error())
	}
}

// collectInventory collects inventory from varnish.params file
func collectInventory(systemEntity *integration.Entity, argList *args2.ArgumentList) {
	if err := collectParamsFile(systemEntity, argList.ParamsConfigFile); err != nil {
		log.Error("Error parsing params file %s: %s", argList.ParamsConfigFile, err.Error())
	}
}

func collectParamsFile(systemEntity *integration.Entity, argsParamLoc string) error {
	// Find if file exists and if so where at
	paramsLoc, err := determineParamsFileLoc(argsParamLoc)
	if err != nil {
		return err
	}

	// Open file
	file, err := os.Open(*paramsLoc)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Warn("Error closing file %s: %s", *paramsLoc, err.Error())
		}
	}()

	// Parse parameters
	params, err := parseParamsFile(file)
	if err != nil {
		return err
	}

	// Set all values as inventory items
	for key, value := range params {
		setInventoryItem(systemEntity, "params/"+key, value)
	}

	return nil
}

// determineParamsFileLoc checks if the params file is present. If
// argsParamLoc is not specify will check in known locations on
// Debian/Ubuntu and RHEL/CentOS systems.
func determineParamsFileLoc(argsParamLoc string) (*string, error) {
	if argsParamLoc != "" {
		if _, err := os.Stat(argsParamLoc); os.IsNotExist(err) {
			return nil, err
		}
		return &argsParamLoc, nil
	}

	// Try Debian/Ubuntu path
	paramsLoc := debianUbuntuParamsLoc
	if _, err := os.Stat(paramsLoc); !os.IsNotExist(err) {
		return &paramsLoc, nil
	}

	// Try RHEL/CentOS
	paramsLoc = rhelCentosParamsLoc
	if _, err := os.Stat(paramsLoc); !os.IsNotExist(err) {
		return &paramsLoc, nil
	}

	return nil, errors.New("varnish.params file could not be found")
}

func parseParamsFile(file *os.File) (map[string]string, error) {
	params := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Comment line
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Param line
		if equal := strings.Index(line, "="); equal > 0 {
			key, value := strings.TrimSpace(line[:equal]), strings.TrimSpace(line[equal+1:])
			// Trim encasing quotes
			value = strings.Trim(value, `"`)
			params[key] = value
		}
	}

	return params, scanner.Err()
}

func setInventoryItem(entity *integration.Entity, key string, value interface{}) {
	if err := entity.SetInventoryItem(key, "value", value); err != nil {
		log.Debug("Error setting Inventory item '%s': %s", key, entity.Metadata.Name)
	}
}
