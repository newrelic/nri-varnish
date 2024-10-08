//go:build integration

package jsonschema

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

// Validate validates the input argument against JSON schema. If the
// input is not valid the error is returned. The first argument is the file name
// of the JSON schema. It is used to build file URI required to load the JSON schema.
// The second argument is the input string that is validated.
func Validate(t *testing.T, fileName string, input string) {
	t.Helper()
	t.Logf("Integration stdout: %s", input)

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	schemaURI := fmt.Sprintf("file://%s", filepath.Join(pwd, fileName))

	schemaLoader := gojsonschema.NewReferenceLoader(schemaURI)
	documentLoader := gojsonschema.NewStringLoader(input)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		t.Errorf("Error loading JSON schema, error: %v", err)
		t.Fail()
	}

	if !result.Valid() {
		t.Errorf("Errors for JSON schema: '%s'\n", schemaURI)
		for _, desc := range result.Errors() {
			t.Errorf("\t- %s\n", desc)
		}

		t.Errorf("The output of the integration doesn't have expected JSON format")
		t.Fail()
	}
}
