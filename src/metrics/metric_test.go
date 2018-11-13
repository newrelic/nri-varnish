package metrics

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"
)

func TestCollectMetrics(t *testing.T) {
	setENV(t, "GO_WANT_HELPER_PROCESS", "1")
	ExecCommand = fakeExecCommand
	defer func() {
		os.Unsetenv("GO_WANT_HELPER_PROCESS")
		ExecCommand = exec.Command
	}()

	i, err := integration.New("test", "1.0.0")
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}

	systemEntity, err := i.Entity("test", "instance")
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}

	if err := CollectMetrics(systemEntity, i); err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	// System Entity Check
	expectedMetrics := []map[string]interface{}{
		{ //main
			"displayName":              systemEntity.Metadata.Name,
			"entityName":               systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name,
			"event_type":               "VarnishSample",
			"mgt.uptimeInMilliseconds": float64(253792),
			"main.threads":             float64(4),
		},
		{ // lock
			"displayName":    systemEntity.Metadata.Name,
			"entityName":     systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name,
			"lock":           "sma",
			"event_type":     "VarnishLockSample",
			"lock.created":   float64(0),
			"lock.destroyed": float64(0),
			"lock.locks":     float64(0),
		},
		{ // storage
			"displayName":                systemEntity.Metadata.Name,
			"entityName":                 systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name,
			"storage":                    "Transient",
			"event_type":                 "VarnishStorageSample",
			"storage.allocReqs":          float64(0),
			"storage.allocFails":         float64(0),
			"storage.allocInBytes":       float64(0),
			"storage.freeInBytes":        float64(0),
			"storage.allocOustanding":    float64(0),
			"storage.outstandingInBytes": float64(40),
			"storage.availableInBytes":   float64(60),
		},
		{ // mempool
			"displayName":                  systemEntity.Metadata.Name,
			"entityName":                   systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name,
			"memoryPool":                   "sess1",
			"event_type":                   "VarnishMempoolSample",
			"mempool.live":                 float64(0),
			"mempool.pool":                 float64(10),
			"mempool.requestSizeInBytes":   float64(512),
			"mempool.allocatedSizeInBytes": float64(480),
			"mempool.allocs":               float64(0),
			"mempool.frees":                float64(0),
			"mempool.recycles":             float64(0),
			"mempool.timeouts":             float64(0),
			"mempool.tooSmall":             float64(0),
			"mempool.surplus":              float64(0),
			"mempool.ranDry":               float64(0),
		},
	}

	for _, set := range systemEntity.Metrics {
		found := false
		for _, expected := range expectedMetrics {
			if reflect.DeepEqual(set.Metrics, expected) {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("Unexpected metric set %+v", set.Metrics)
		}
	}

	// Backend Entity Check
	backendEntity := findBackendEntity(i)
	if backendEntity == nil {
		t.Fatalf("No backends found")
	}

	expectedBackendMetrics := map[string]interface{}{
		"displayName":                       "boot.web1",
		"entityName":                        "backend:boot.web1",
		"event_type":                        "VarnishBackendSample",
		"varnishInstance":                   systemEntity.Metadata.Name,
		"backend.happy":                     float64(0),
		"net.backend.requestHeaderInBytes":  float64(0),
		"net.backend.requestBodyInBytes":    float64(0),
		"net.backend.responseHeaderInBytes": float64(0),
		"net.backend.responseBodyInBytes":   float64(0),
		"net.backend.pipeHeaderInBytes":     float64(0),
		"net.backend.pipeOutInBytes":        float64(0),
		"net.backend.pipeInInBytes":         float64(0),
		"backend.connections":               float64(0),
		"net.backend.requests":              float64(0),
	}

	backendResult := backendEntity.Metrics[0].Metrics
	if !reflect.DeepEqual(backendResult, expectedBackendMetrics) {
		t.Errorf("Expected backend %+v got %+v", expectedBackendMetrics, backendResult)
	}
}

func findBackendEntity(i *integration.Integration) *integration.Entity {
	for _, entity := range i.Entities {
		if entity.Metadata.Namespace == "backend" {
			return entity
		}
	}

	return nil
}

func TestCollectMetrics_ExecError(t *testing.T) {
	setENV(t, "GO_WANT_HELPER_PROCESS", "1")
	setENV(t, "GO_WANT_ERROR", "1")
	ExecCommand = fakeExecCommand

	defer func() {
		ExecCommand = exec.Command
		os.Unsetenv("GO_WANT_ERROR")
		os.Unsetenv("GO_WANT_HELPER_PROCESS")
	}()

	i, err := integration.New("test", "1.0.0")
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}

	systemEntity, err := i.Entity("test", "varnish")
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}

	if err := CollectMetrics(systemEntity, i); err == nil {
		t.Error("Expected error")
	}
}

func setENV(t *testing.T, envName, value string) {
	if err := os.Setenv(envName, value); err != nil {
		t.Fatalf("Failed to set env %s: %s", envName, err.Error())
	}
}

func fakeExecCommand(command string, args ...string) (cmd *exec.Cmd) {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd = exec.Command(os.Args[0], cs...)
	return cmd
}

// TestHelperProcess isn't a real test. It's used as a helper process.
func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	if os.Getenv("GO_WANT_ERROR") == "1" {
		os.Exit(2)
	} else {
		fmt.Fprintf(os.Stdout, varnishStatTestResult)
		os.Exit(0)
	}
}
