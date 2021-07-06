package metrics

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/log"
)

func parseStats(statsData []byte) (*varnishDefinition, map[string]*backendDefinition, error) {
	stats := make(map[string]interface{})
	if err := json.Unmarshal(statsData, &stats); err != nil {
		return nil, nil, err
	}

	varnishSystem := &varnishDefinition{
		locks:    make(map[string]*lockDefinition),
		mempools: make(map[string]*mempoolDefinition),
		storages: make(map[string]*storageDefinition),
	}

	backends := make(map[string]*backendDefinition)

	version, ok := stats["version"].(float64)
	if !ok {
		log.Debug("No 'version' key found, assuming legacy varnishstat format")
		version = 0
	}

	if version < 0 || version > 1 {
		log.Debug("Unknown varnishstat results version: %d", version)
		return varnishSystem, backends, nil
	}

	for key, entry := range stats {
		if key == "timestamp" || key == "version" {
			// skip the timestamp and version fields
			continue
		}

		if version == 0 {
			processStat(key, entry, varnishSystem, backends)
		} else if version == 1 {
			if key == "counters" {
				counters, ok := entry.(map[string]interface{})
				if !ok {
					log.Debug("Error parsing 'counters': non object map type %T", entry)
					continue
				}
				for fullStatName, statEntry := range counters {
					processStat(fullStatName, statEntry, varnishSystem, backends)
				}
			}
		}
	}

	return varnishSystem, backends, nil
}

func processStat(fullStatName string, entry interface{}, varnishSystem *varnishDefinition, backends map[string]*backendDefinition) {
	// retrieve value
	statValue, err := getStatValue(entry)
	if err != nil {
		log.Debug("Error parsing metric '%s': %s", fullStatName, err.Error())
		return
	}

	if strings.HasPrefix(fullStatName, "VBE.") {
		// backends
		setBackendValue(backends, fullStatName, statValue)
	} else {
		// all other metrics under varnish system
		parseAndSetStat(varnishSystem, fullStatName, statValue)
	}
}

// getStatValue retrieves the "value" field from the underlying map
func getStatValue(stat interface{}) (interface{}, error) {
	statEntry, ok := stat.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("non object map type %T", stat)
	}

	statValue, ok := statEntry["value"]
	if !ok {
		return nil, errors.New("no value found")
	}

	return statValue, nil
}

func parseAndSetStat(varnishSystem *varnishDefinition, fullStatName string, statValue interface{}) {
	// Parse out into structs
	if strings.HasPrefix(fullStatName, "LCK.") {
		// locks
		setLockValue(varnishSystem.locks, fullStatName, statValue)
	} else if strings.HasPrefix(fullStatName, "SMF.") || strings.HasPrefix(fullStatName, "SMU.") || strings.HasPrefix(fullStatName, "SMA.") {
		// storage
		setStorageValue(varnishSystem.storages, fullStatName, statValue)
	} else if strings.HasPrefix(fullStatName, "MEMPOOL.") {
		// mempools
		setMempoolValue(varnishSystem.mempools, fullStatName, statValue)
	} else {
		// main sample
		setSystemValue(varnishSystem, fullStatName, statValue)
	}
}

// The following setXValue functions are just repeated code but due to Go not having generics it's
// cleaner than doing some nasty reflection in order to achieve a generic solution

func setSystemValue(varnishSystem *varnishDefinition, fullStatName string, statValue interface{}) {
	_, statName := parseStatName(fullStatName)

	if err := setValue(varnishSystem, statName, statValue); err != nil {
		log.Debug("Error setting metric value for stat '%s': %s", fullStatName, err.Error())
	}
}

func setLockValue(lockMap map[string]*lockDefinition, fullStatName string, statValue interface{}) {
	lockName, statName := parseStatName(fullStatName)
	lock, ok := lockMap[lockName]
	if !ok {
		lock = new(lockDefinition)
		lockMap[lockName] = lock
	}

	if err := setValue(lock, statName, statValue); err != nil {
		log.Debug("Error setting metric value for stat '%s' on lock '%s': %s", statName, lockName, err.Error())
	}
}

func setBackendValue(backendMap map[string]*backendDefinition, fullStatName string, statValue interface{}) {
	backendName, statName := parseStatName(fullStatName)
	backend, ok := backendMap[backendName]
	if !ok {
		backend = new(backendDefinition)
		backendMap[backendName] = backend
	}

	if err := setValue(backend, statName, statValue); err != nil {
		log.Debug("Error setting metric value for stat '%s' on backend '%s': %s", statName, backendName, err.Error())
	}
}

func setStorageValue(storageMap map[string]*storageDefinition, fullStatName string, statValue interface{}) {
	storageName, statName := parseStatName(fullStatName)
	storage, ok := storageMap[storageName]
	if !ok {
		storage = new(storageDefinition)
		storageMap[storageName] = storage
	}

	if err := setValue(storage, statName, statValue); err != nil {
		log.Debug("Error setting metric value for stat '%s' on storage '%s': %s", statName, storageName, err.Error())
	}
}

func setMempoolValue(mempoolMap map[string]*mempoolDefinition, fullStatName string, statValue interface{}) {
	mempoolName, statName := parseStatName(fullStatName)
	mempool, ok := mempoolMap[mempoolName]
	if !ok {
		mempool = new(mempoolDefinition)
		mempoolMap[mempoolName] = mempool
	}

	if err := setValue(mempool, statName, statValue); err != nil {
		log.Debug("Error setting metric value for stat '%s' on mempool '%s': %s", statName, mempoolName, err.Error())
	}
}

// parseStatName parses out a full stat name of the form PREFIX.ENTITY_NAME.STAT
// into just the entity name and the stat name
func parseStatName(fullStatName string) (entityName, statname string) {
	stringParts := strings.Split(fullStatName, ".")
	statname = stringParts[len(stringParts)-1]

	// special case for uptime
	if statname == "uptime" {
		statname = stringParts[0] + "." + statname
	}

	// When the fullStatName contains no prefixes, just return it
	if len(stringParts) == 1 {
		return
	}

	// Get middle chunk of strings
	entityParts := stringParts[1 : len(stringParts)-1]
	entityName = strings.Join(entityParts, ".")
	return
}

// setValue sets the appropriate struct field via reflection by matching the
// struct tag 'stat_name' with the stat name from the Varnish stat.
// If not match is found then an error is returned
func setValue(v interface{}, statName string, metricValue interface{}) error {
	value := reflect.ValueOf(v).Elem()
	t := value.Type()

	statValue := reflect.ValueOf(metricValue)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag, hasStatName := f.Tag.Lookup("stat_name")
		if !hasStatName || tag != statName {
			continue
		}

		fv := value.FieldByName(f.Name)
		fv.Set(statValue)
		return nil
	}

	return fmt.Errorf("no field found for stat '%s'", statName)
}
