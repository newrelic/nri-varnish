package main

import (
	"reflect"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/data/inventory"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-varnish/internal/args"
)

func TestCollectInventory_Normal(t *testing.T) {
	i, err := integration.New("test", "1.0.0")
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}

	entity, err := i.Entity("test", "varnish")
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}

	argList := &args.ArgumentList{
		ParamsConfigFile: "testdata/varnish.params",
	}

	expected := inventory.Items{
		"params/RELOAD_VCL": inventory.Item{
			"value": "1",
		},
		"params/VARNISH_VCL_CONF": inventory.Item{
			"value": "/etc/varnish/default.vcl",
		},
		"params/VARNISH_LISTEN_PORT": inventory.Item{
			"value": "6081",
		},
		"params/VARNISH_ADMIN_LISTEN_ADDRESS": inventory.Item{
			"value": "0.0.0.0",
		},
		"params/VARNISH_ADMIN_LISTEN_PORT": inventory.Item{
			"value": "6082",
		},
		"params/VARNISH_SECRET_FILE": inventory.Item{
			"value": "/etc/varnish/secret",
		},
		"params/VARNISH_STORAGE": inventory.Item{
			"value": "malloc,256M",
		},
		"params/VARNISH_USER": inventory.Item{
			"value": "varnish",
		},
		"params/VARNISH_GROUP": inventory.Item{
			"value": "varnish",
		},
	}

	collectInventory(entity, argList)

	out := entity.Inventory.Items()
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Expected %+v got %+v", expected, out)
	}
}

func TestCollectInventory_NoParamsFile(t *testing.T) {
	i, err := integration.New("test", "1.0.0")
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}

	entity := i.LocalEntity()

	argList := &args.ArgumentList{}

	expected := inventory.Items{}

	collectInventory(entity, argList)

	out := entity.Inventory.Items()
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Expected %+v got %+v", expected, out)
	}
}
