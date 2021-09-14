//go:build integration
// +build integration

package integration

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/newrelic/nri-varnish/tests/integration/helpers"
	"github.com/newrelic/nri-varnish/tests/integration/jsonschema"
	"github.com/stretchr/testify/assert"
)

const (
	defaultContainer     = "integration_nri-varnish_1"
	defaultBinPath       = "/nri-varnish"
	schemasPath          = "json-schema-files"
	metricSchemasFile    = "metrics.json"
	inventorySchemasFile = "inventory.json"
)

// Returns the standard output, or fails testing if the command returned an error
func runIntegration(t *testing.T, envVars ...string) (string, string) {
	t.Helper()

	command := make([]string, 0)
	command = append(command, defaultBinPath)

	timeout := 10 * time.Second

	stdout, stderr := helpers.ExecInContainer(t, timeout, defaultContainer, command, envVars...)

	if stderr != "" {
		t.Logf("Integration command Standard Error: %s", stderr)
	}

	return stdout, stderr
}

func Test_integration_collects_metrics(t *testing.T) {
	t.Parallel()

	stdout, stderr := runIntegration(t, "INSTANCE_NAME=test", "METRICS=true")
	assert.Empty(t, stderr)

	varnishSchemaPath := filepath.Join(schemasPath, metricSchemasFile)
	jsonschema.Validate(t, varnishSchemaPath, stdout)
}

func Test_integration_collects_inventory(t *testing.T) {
	t.Parallel()

	stdout, stderr := runIntegration(t, "INSTANCE_NAME=test", "INVENTORY=true", "PARAMS_CONFIG_FILE=/testdata/varnish.params")
	assert.Empty(t, stderr)

	varnishSchemaPath := filepath.Join(schemasPath, inventorySchemasFile)
	jsonschema.Validate(t, varnishSchemaPath, stdout)
}

func Test_integration_not_able_to_find_params(t *testing.T) {
	t.Parallel()

	stdout, stderr := runIntegration(t, "INSTANCE_NAME=test")
	assert.Contains(t, stderr, "Error parsing params file")

	varnishSchemaPath := filepath.Join(schemasPath, metricSchemasFile)
	jsonschema.Validate(t, varnishSchemaPath, stdout)
}
