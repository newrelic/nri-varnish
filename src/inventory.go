package main

import (
	"strings"
	"bufio"
	"os"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-varnish/src/args"
	"github.com/newrelic/infra-integrations-sdk/log"
)

// CollectInventory collects inventory from 
func CollectInventory(entity *integration.Entity, argList *args.ArgumentList) {
	if err := collectParamsFile(entity, argList.ParamsConfigFile); err != nil {
		log.Error("Error parsing params file '%s': %s", argList.ParamsConfigFile, err.Error())
	}


}

func collectParamsFile(entity *integration.Entity, paramsLoc string) error {
	

	file, err := os.Open(paramsLoc)
	if err != nil {
		return err
	}
	defer file.Close()

	params, err := parseParamsFile(file)
	if err != nil {
		return err
	}

	// Set all values as inventory items
	for key, value := range params {
		setInventoryItem(entity, "params/" + key, value)
	}

	return nil
}

func determineParamsFileLoc() (string, error) {
	if _, err := os.Stat(paramsLoc); os.IsNotExist(err) {
		return err
	}
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