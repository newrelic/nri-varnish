// Package metrics handles metric collection from Varnish
package metrics

import (
	"os/exec"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

var execCommand = exec.Command

// CollectMetrics collects metrics from varnishstat command
func CollectMetrics(systemEntity *integration.Entity, i *integration.Integration) error {
	statOutput, err := execCommand("varnishstat", "-j").Output()
	if err != nil {
		log.Error("Error running varnishstat command: %s", err.Error())
		return err
	}

	varnishSystem, backends, err := parseStats(statOutput)
	if err != nil {
		log.Error("Error parsing varnishstat output: %s", err.Error())
		return err
	}

	// create backend entities and set metrics
	processBackends(backends, i)

	// process metrics for varnishSystem
	processVarnishSystem(systemEntity, varnishSystem)

	return nil
}

func processBackends(backends map[string]*backendDefinition, i *integration.Integration) {
	for backendName, def := range backends {
		entity, err := i.Entity(backendName, "backend")
		if err != nil {
			log.Debug("Error creating entity for Backend %s: %s", backendName, err.Error())
			continue
		}

		metricSet := entity.NewMetricSet("VarnishBackendSample",
			metric.Attribute{Key: "displayName", Value: entity.Metadata.Name},
			metric.Attribute{Key: "entityName", Value: entity.Metadata.Namespace + ":" + entity.Metadata.Name},
		)

		if err := metricSet.MarshalMetrics(def); err != nil {
			log.Warn("Error setting metrics for backend %s: %s", backendName, err.Error())
			continue
		}
	}
}

func processVarnishSystem(systemEntity *integration.Entity, varnishSystem *varnishDefinition) {
	metricSet := systemEntity.NewMetricSet("VarnishSample",
		metric.Attribute{Key: "displayName", Value: systemEntity.Metadata.Name},
		metric.Attribute{Key: "entityName", Value: systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name},
	)

	if err := metricSet.MarshalMetrics(varnishSystem); err != nil {
		log.Warn("Error setting metrics for Varnish System: %s", err.Error())
	}

	// Process lock samples
	for lockName, lock := range varnishSystem.locks {
		if err := processSubSample(lock, "VarnishLockSample", systemEntity); err != nil {
			log.Warn("Error setting metrics for Lock %s: %s", lockName, err.Error())
		}
	}

	// Process mempool samples
	for mempoolName, mempool := range varnishSystem.mempools {
		if err := processSubSample(mempool, "VarnishMempoolSample", systemEntity); err != nil {
			log.Warn("Error setting metrics for Mempool %s: %s", mempoolName, err.Error())
		}
	}

	// Process storage samples
	for storageName, storage := range varnishSystem.storages {
		if err := processSubSample(storage, "VarnishStorageSample", systemEntity); err != nil {
			log.Warn("Error setting metrics for Storage %s: %s", storageName, err.Error())
		}
	}
}

func processSubSample(subStructure interface{}, sampleName string, systemEntity *integration.Entity) error {
	metricSet := systemEntity.NewMetricSet(sampleName,
		metric.Attribute{Key: "displayName", Value: systemEntity.Metadata.Name},
		metric.Attribute{Key: "entityName", Value: systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name},
	)

	return metricSet.MarshalMetrics(subStructure)
}
