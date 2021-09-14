package metrics

import (
	"reflect"
	"testing"
)

const varnishStatTestResult = `{
	"timestamp": "2018-11-05T17:22:34",
	"backend_fail": {
		"value": 0,
		"flag": "a",
		"description": "Backend conn. failures"
	},
	"MGT.uptime": {
		"description": "Management process uptime",
		"flag": "c", "format": "d",
		"value": 253792
	},
	"MAIN.threads": {
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
		"value": 40
	},
	"SMA.Transient.g_space": {
		"description": "Bytes available",
		"flag": "g", "format": "B",
		"value": 60
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

func Test_parseStats_Full(t *testing.T) {
	expectedVarnishSystem := &varnishDefinition{
		MgtUptime:       float64(253792),
		MainThreads:     float64(4),
		BackendConFails: float64(0),

		locks: map[string]*lockDefinition{
			"sma": {
				Created:   float64(2),
				Destroyed: float64(0),
				LockOps:   float64(4),
			},
		},
		storages: map[string]*storageDefinition{
			"Transient": {
				AllocReqs:       float64(0),
				AllocFails:      float64(0),
				Alloc:           float64(0),
				Free:            float64(0),
				AllocOustanding: float64(0),
				Outstanding:     float64(40),
				Available:       float64(60),
			},
		},
		mempools: map[string]*mempoolDefinition{
			"sess1": {
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

	varnishSystem, backends, err := parseStats([]byte(varnishStatTestResult))
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

const varnishStatV1TestResult = `{
	"version": 1,
	"timestamp": "2021-06-23T15:25:46",
	"counters": {
		"MGT.uptime": {
			"description": "Management process uptime",
			"flag": "c",
			"format": "d",
			"value": 4173714
		},
		"MAIN.threads": {
			"description": "Total number of threads",
			"flag": "g",
			"format": "i",
			"value": 200
		},
		"LCK.backend.creat": {
			"description": "Created locks",
			"flag": "c",
			"format": "i",
			"value": 2
		},
		"LCK.backend.destroy": {
			"description": "Destroyed locks",
			"flag": "c",
			"format": "i",
			"value": 0
		},
		"LCK.backend.locks": {
			"description": "Lock Operations",
			"flag": "c",
			"format": "i",
			"value": 94958
		},
		"SMA.Transient.g_bytes": {
			"description": "Bytes outstanding",
			"flag": "g", "format": "B",
			"value": 40
		},
		"SMA.Transient.g_space": {
			"description": "Bytes available",
			"flag": "g", "format": "B",
			"value": 60
		},
		"MEMPOOL.busyobj.live": {
			"description": "In use",
			"flag": "g",
			"format": "i",
			"value": 0
		},
		"MEMPOOL.busyobj.pool": {
			"description": "In Pool",
			"flag": "g",
			"format": "i",
			"value": 10
		},
		"VBE.reload_20210601_145132_6770.upstream_1.fail_eacces": {
			"description": "Connections failed with EACCES or EPERM",
			"flag": "c",
			"format": "i",
			"value": 0
		},
		"VBE.reload_20210601_145132_6770.upstream_1.fail_eaddrnotavail": {
			"description": "Connections failed with EADDRNOTAVAIL",
			"flag": "c",
			"format": "i",
			"value": 0
		},
		"VBE.reload_20210601_145132_6770.upstream_1.fail_econnrefused": {
			"description": "Connections failed with ECONNREFUSED",
			"flag": "c",
			"format": "i",
			"value": 0
		},
		"VBE.reload_20210601_145132_6770.upstream_1.fail_enetunreach": {
			"description": "Connections failed with ENETUNREACH",
			"flag": "c",
			"format": "i",
			"value": 0
		},
		"VBE.reload_20210601_145132_6770.upstream_1.fail_etimedout": {
			"description": "Connections failed ETIMEDOUT",
			"flag": "c",
			"format": "i",
			"value": 0
		}
	}
}`

func Test_parseStats_uknownVersion(t *testing.T) {
	testCases := []struct {
		name       string
		stats      string
		expectFail bool
	}{
		{
			"empty ok",
			`{ "version": 0 }`,
			false,
		},
		{
			"bad json",
			`{ "version": 0, }`,
			true,
		},
		{
			"unknown schema version",
			`{ "version": 9.1 }`,
			true,
		},
		{
			"string version fail to cast",
			`{ "version": "asdf" }`,
			true,
		},
		{
			"v1 missing counters entry",
			`{
				"version": 1,
				"timestamp": "2021-06-23T15:25:46",
				"MGT.uptime": {
					"description": "Management process uptime",
					"flag": "c",
					"format": "d",
					"value": 4173714
				}
			}`,
			true,
		},
	}

	for _, tc := range testCases {
		_, _, err := parseStats([]byte(tc.stats))
		if tc.expectFail && err == nil {
			t.Errorf("Test Case %s Failed: did not return error", tc.name)
		}
		if !tc.expectFail && err != nil {
			t.Errorf("Test Case %s Failed: return error: %s", tc.name, err)
		}
	}
}

func Test_parseStats_FullV1(t *testing.T) {
	expectedVarnishSystem := &varnishDefinition{
		MgtUptime:   float64(4173714),
		MainThreads: float64(200),

		locks: map[string]*lockDefinition{
			"backend": {
				Created:   float64(2),
				Destroyed: float64(0),
				LockOps:   float64(94958),
			},
		},
		storages: map[string]*storageDefinition{
			"Transient": {
				Outstanding: float64(40),
				Available:   float64(60),
			},
		},
		mempools: map[string]*mempoolDefinition{
			"busyobj": {
				Live: float64(0),
				Pool: float64(10),
			},
		},
	}

	expectedBackends := map[string]*backendDefinition{
		"reload_20210601_145132_6770.upstream_1": {
			ConFailedEaccess:       float64(0),
			ConFailedEaddrnotavail: float64(0),
			ConFailedEconnrefused:  float64(0),
			ConFailedEnetunreach:   float64(0),
			ConFailedEtimedout:     float64(0),
		},
	}

	varnishSystem, backends, err := parseStats([]byte(varnishStatV1TestResult))
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
