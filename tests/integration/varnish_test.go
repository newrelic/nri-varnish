//go:build integration

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
	varnishStatsV0                = "varnish-stats-v0"
	varnishStatsV1                = "varnish-stats-v1"
	varnishStatsNamedInstance     = "varnish-named-instance"
	varnishStatsNamedInstanceName = "customInstanceName" // name defined in the container commands
	defaultBinPath                = "/nri-varnish"
	schemasPath                   = "json-schema-files"
	metricSchemasFile             = "metrics.json"
	inventorySchemasFile          = "inventory.json"
)

func Test_integration_test(t *testing.T) {
	t.Parallel()

	containers := []string{varnishStatsV0, varnishStatsV1}

	testCases := []struct {
		name       string
		envs       []string
		schemaFile string
	}{
		{
			"collects metrics only",
			[]string{"METRICS=true"},
			metricSchemasFile,
		},
		{
			"collects inventory only",
			[]string{"INVENTORY=true", "PARAMS_CONFIG_FILE=/testdata/varnish.params"},
			inventorySchemasFile,
		},
	}

	for _, tc := range testCases {
		tc.envs = append(tc.envs, "INSTANCE_NAME=test")
		for _, container := range containers {
			stdout, _, err := runIntegration(t, container, tc.envs...)
			assert.NoError(t, err)

			varnishSchemaPath := filepath.Join(schemasPath, tc.schemaFile)
			jsonschema.Validate(t, varnishSchemaPath, stdout)
		}
	}
}

func Test_integration_not_able_to_find_params(t *testing.T) {
	t.Parallel()

	_, stderr, err := runIntegration(t, varnishStatsV0, "INSTANCE_NAME=test", "INVENTORY=true")
	assert.NoError(t, err)
	assert.Contains(t, stderr, "Error parsing params file")
}

func Test_named_instance(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		envs []string
		fail bool
	}{
		{
			"metric collection fails without defined instance name",
			[]string{"METRICS=true"},
			true,
		},
		{
			"metric collection works with defined instance name",
			[]string{"METRICS=true", "VARNISH_NAME=" + varnishStatsNamedInstanceName},
			false,
		},
	}

	for _, tc := range testCases {
		tc.envs = append(tc.envs, "INSTANCE_NAME=test")

		stdout, _, err := runIntegration(t, varnishStatsNamedInstance, tc.envs...)
		if tc.fail {
			assert.Error(t, err)
			continue
		}

		assert.NoError(t, err)
		varnishSchemaPath := filepath.Join(schemasPath, metricSchemasFile)
		jsonschema.Validate(t, varnishSchemaPath, stdout)
	}
}

// Returns the standard output, or fails testing if the command returned an error
func runIntegration(t *testing.T, container string, envVars ...string) (string, string, error) {
	t.Helper()

	command := make([]string, 0)
	command = append(command, defaultBinPath)

	timeout := 10 * time.Second

	stdout, stderr, err := helpers.ExecInContainer(t, timeout, container, command, envVars...)

	if stderr != "" {
		t.Logf("Integration command Standard Error: %s", stderr)
	}

	return stdout, stderr, err
}
