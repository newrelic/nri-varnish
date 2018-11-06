package metrics

import (
	"reflect"
	"testing"
)

func Test_parseStats_Full(t *testing.T) {
	input := `{
	"timestamp": "2018-11-05T17:22:34",
	"MGT.uptime": {
		"description": "Management process uptime",
		"flag": "c", "format": "d",
		"value": 253792
	},
	"MAIN.summs": {
		"description": "stat summ operations",
		"flag": "c", "format": "i",
		"value": 4
	},
	"LCK.sma.creat": {
		"description": "Created locks",
		"flag": "c", "format": "i",
		"value": 2
	},
	"LCK.sma.destroy": {
		"description": "Destroyed locks",
		"flag": "c", "format": "i",
		"value": 0
	},
	"LCK.sma.locks": {
		"description": "Lock Operations",
		"flag": "c", "format": "i",
		"value": 4
	},
	"SMA.Transient.c_req": {
		"description": "Allocator requests",
		"flag": "c", "format": "i",
		"value": 0
	},
	"SMA.Transient.c_fail": {
		"description": "Allocator failures",
		"flag": "c", "format": "i",
		"value": 0
	},
	"SMA.Transient.c_bytes": {
		"description": "Bytes allocated",
		"flag": "c", "format": "B",
		"value": 0
	},
	"SMA.Transient.c_freed": {
		"description": "Bytes freed",
		"flag": "c", "format": "B",
		"value": 0
	},
	"SMA.Transient.g_alloc": {
		"description": "Allocations outstanding",
		"flag": "g", "format": "i",
		"value": 0
	},
	"SMA.Transient.g_bytes": {
		"description": "Bytes outstanding",
		"flag": "g", "format": "B",
		"value": 0
	},
	"SMA.Transient.g_space": {
		"description": "Bytes available",
		"flag": "g", "format": "B",
		"value": 0
	},
	"MEMPOOL.sess1.live": {
		"description": "In use",
		"flag": "g", "format": "i",
		"value": 0
	},
	"MEMPOOL.sess1.pool": {
		"description": "In Pool",
		"flag": "g", "format": "i",
		"value": 10
	},
	"MEMPOOL.sess1.sz_wanted": {
		"description": "Size requested",
		"flag": "g", "format": "B",
		"value": 512
	},
	"MEMPOOL.sess1.sz_actual": {
		"description": "Size allocated",
		"flag": "g", "format": "B",
		"value": 480
	},
	"MEMPOOL.sess1.allocs": {
		"description": "Allocations",
		"flag": "c", "format": "i",
		"value": 0
	},
	"MEMPOOL.sess1.frees": {
		"description": "Frees",
		"flag": "c", "format": "i",
		"value": 0
	},
	"MEMPOOL.sess1.recycle": {
		"description": "Recycled from pool",
		"flag": "c", "format": "i",
		"value": 0
	},
	"MEMPOOL.sess1.timeout": {
		"description": "Timed out from pool",
		"flag": "c", "format": "i",
		"value": 0
	},
	"MEMPOOL.sess1.toosmall": {
		"description": "Too small to recycle",
		"flag": "c", "format": "i",
		"value": 0
	},
	"MEMPOOL.sess1.surplus": {
		"description": "Too many for pool",
		"flag": "c", "format": "i",
		"value": 0
	},
	"MEMPOOL.sess1.randry": {
		"description": "Pool ran dry",
		"flag": "c", "format": "i",
		"value": 0
	},
	"VBE.boot.web1.happy": {
		"description": "Happy health probes",
		"flag": "b", "format": "b",
		"value": 0
	},
	"VBE.boot.web1.bereq_hdrbytes": {
		"description": "Request header bytes",
		"flag": "c", "format": "B",
		"value": 381
	},
	"VBE.boot.web1.bereq_bodybytes": {
		"description": "Request body bytes",
		"flag": "c", "format": "B",
		"value": 0
	},
	"VBE.boot.web1.beresp_hdrbytes": {
		"description": "Response header bytes",
		"flag": "c", "format": "B",
		"value": 229
	},
	"VBE.boot.web1.beresp_bodybytes": {
		"description": "Response body bytes",
		"flag": "c", "format": "B",
		"value": 2326
	},
	"VBE.boot.web1.pipe_hdrbytes": {
		"description": "Pipe request header bytes",
		"flag": "c", "format": "B",
		"value": 0
	},
	"VBE.boot.web1.pipe_out": {
		"description": "Piped bytes to backend",
		"flag": "c", "format": "B",
		"value": 0
	},
	"VBE.boot.web1.pipe_in": {
		"description": "Piped bytes from backend",
		"flag": "c", "format": "B",
		"value": 0
	},
	"VBE.boot.web1.conn": {
		"description": "Concurrent connections to backend",
		"flag": "g", "format": "i",
		"value": 0
	},
	"VBE.boot.web1.req": {
		"description": "Backend requests sent",
		"flag": "c", "format": "i",
		"value": 1
	}
}`

	expectedVarnishSystem := &varnishDefinition{
		MgtUptime: float64(253792),
		MainSumms: float64(4),

		locks: map[string]*lockDefinition{
			"sma": {
				Name:      "sma",
				Created:   float64(2),
				Destroyed: float64(0),
				LockOps:   float64(4),
			},
		},
		storages: map[string]*storageDefinition{
			"Transient": {
				Name:            "Transient",
				AllocReqs:       float64(0),
				AllocFails:      float64(0),
				Alloc:           float64(0),
				Free:            float64(0),
				AllocOustanding: float64(0),
				Outstanding:     float64(0),
				Available:       float64(0),
			},
		},
		mempools: map[string]*mempoolDefinition{
			"sess1": {
				Name:          "sess1",
				Live:          float64(0),
				Pool:          float64(10),
				ReqSize:       float64(512),
				AllocatedSize: float64(480),
				Allocs:        float64(0),
				Frees:         float64(0),
				Recycles:      float64(0),
				Timeouts:      float64(0),
				TooSmall:      float64(0),
				Surplus:       float64(0),
				RanDry:        float64(0),
			},
		},
	}

	expectedBackends := map[string]*backendDefinition{
		"boot.web1": {
			Happy:       float64(0),
			ReqHeader:   float64(381),
			ReqBody:     float64(0),
			RespHeader:  float64(229),
			RespBody:    float64(2326),
			PipeHeader:  float64(0),
			PipeOut:     float64(0),
			PipeIn:      float64(0),
			Connections: float64(0),
			Req:         float64(1),
		},
	}

	varnishSystem, backends, err := parseStats([]byte(input))
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}

	if !reflect.DeepEqual(varnishSystem, expectedVarnishSystem) {
		t.Fatalf("Expected System of %+v got %+v", expectedVarnishSystem, varnishSystem)
	}

	if !reflect.DeepEqual(backends, expectedBackends) {
		t.Fatalf("Expected Backends of %+v got %+v", expectedBackends, backends)
	}
}

func Test_getStatValue_Errors(t *testing.T) {
	testCases := []struct {
		name   string
		intput interface{}
	}{
		{
			"Bad Map Parse",
			"nope",
		},
		{
			"Missing Value",
			map[string]interface{}{
				"notValue": float64(0),
			},
		},
	}

	for _, tc := range testCases {
		if _, err := getStatValue(tc.intput); err == nil {
			t.Errorf("Test Case %s Failed: did not return error", tc.name)
		}
	}
}

func Test_setValue_NoField(t *testing.T) {
	backend := new(backendDefinition)

	statName, metricValue := "not_a_stat", float64(4)

	if err := setValue(backend, statName, metricValue); err == nil {
		t.Error("Expected error")
	}
}
