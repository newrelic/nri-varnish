// Package metrics handles metric collection from Varnish
package metrics

import (
	"os/exec"

	"github.com/newrelic/infra-integrations-sdk/data/attribute"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

// ExecCommand holds the exec.Command function. Meant
// to be overridden for testing
var ExecCommand = exec.Command

// CollectMetrics collects metrics from varnishstat command
func CollectMetrics(systemEntity *integration.Entity, i *integration.Integration, varnishName string) error {
	argList := []string{"-j"}

	if varnishName != "" {
		argList = append(argList, "-n", varnishName)
	}

	statOutput, err := ExecCommand("varnishstat", argList...).Output()
	if err != nil {
		log.Debug("Error running varnishstat command: %s", err.Error())
		return err
	}

	varnishSystem, backends, err := parseStats(statOutput)
	if err != nil {
		log.Debug("Error parsing varnishstat output: %s", err.Error())
		return err
	}

	// create backend entities and set metrics
	processBackends(systemEntity.Metadata.Name, backends, i)

	// process metrics for varnishSystem
	processVarnishSystem(systemEntity, varnishSystem)

	return nil
}

func processBackends(instanceName string, backends map[string]*backendDefinition, i *integration.Integration) {
	for backendName, def := range backends {
		entityIDAttr := integration.IDAttribute{Key: "instanceName", Value: instanceName}
		entity, err := i.EntityReportedBy("localhost", backendName, "va-backend", entityIDAttr)
		if err != nil {
			log.Debug("Error creating entity for Backend %s: %s", backendName, err.Error())
			continue
		}

		metricSet := entity.NewMetricSet("VarnishBackendSample",
			attribute.Attribute{Key: "displayName", Value: entity.Metadata.Name},
			attribute.Attribute{Key: "entityName", Value: entity.Metadata.Namespace + ":" + entity.Metadata.Name},
		)

		if err := metricSet.MarshalMetrics(def); err != nil {
			log.Warn("Error setting metrics for backend %s: %s", backendName, err.Error())
			continue
		}
	}
}

func processVarnishSystem(systemEntity *integration.Entity, varnishSystem *varnishDefinition) {
	metricSet := systemEntity.NewMetricSet("VarnishSample",
		attribute.Attribute{Key: "displayName", Value: systemEntity.Metadata.Name},
		attribute.Attribute{Key: "entityName", Value: systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name},
	)

	if err := metricSet.MarshalMetrics(varnishSystem); err != nil {
		log.Warn("Error setting metrics for Varnish System: %s", err.Error())
	}

	// Process lock samples
	for lockName, lock := range varnishSystem.locks {
		if err := processSubSample(lock, "VarnishLockSample", "lock", lockName, systemEntity); err != nil {
			log.Warn("Error setting metrics for Lock %s: %s", lockName, err.Error())
		}
	}

	// Process mempool samples
	for mempoolName, mempool := range varnishSystem.mempools {
		if err := processSubSample(mempool, "VarnishMempoolSample", "memoryPool", mempoolName, systemEntity); err != nil {
			log.Warn("Error setting metrics for Mempool %s: %s", mempoolName, err.Error())
		}
	}

	// Process storage samples
	for storageName, storage := range varnishSystem.storages {
		if err := processSubSample(storage, "VarnishStorageSample", "storage", storageName, systemEntity); err != nil {
			log.Warn("Error setting metrics for Storage %s: %s", storageName, err.Error())
		}
	}

	// Process book samples
	for bookName, book := range varnishSystem.book {
		if err := processSubSample(book, "VarnishBookSample", "book", bookName, systemEntity); err != nil {
			log.Warn("Error setting metrics for Book %s: %s", bookName, err.Error())
		}
	}

	// Process store samples
	for storeName, store := range varnishSystem.store {
		if err := processSubSample(store, "VarnishStoreSample", "store", storeName, systemEntity); err != nil {
			log.Warn("Error setting metrics for Store %s: %s", storeName, err.Error())
		}
	}

	// Process massive storage samples
	for massiveStorageName, massiveStorage := range varnishSystem.massiveStorage {
		if err := processSubSample(massiveStorage, "VarnishMassiveStorageSample", "massiveStorage", massiveStorageName, systemEntity); err != nil {
			log.Warn("Error setting metrics for Massive Storage %s: %s", massiveStorageName, err.Error())
		}
	}
}

func processSubSample(subStructure interface{}, sampleName, idAttribute, idAttributeValue string, systemEntity *integration.Entity) error {
	metricSet := systemEntity.NewMetricSet(sampleName,
		attribute.Attribute{Key: "displayName", Value: systemEntity.Metadata.Name},
		attribute.Attribute{Key: "entityName", Value: systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name},
		attribute.Attribute{Key: idAttribute, Value: idAttributeValue},
	)

	return metricSet.MarshalMetrics(subStructure)
}
