// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build linux

package login

import (
	"encoding/binary"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/auditbeat/core"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/paths"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

func TestData(t *testing.T) {
	if byteOrder != binary.LittleEndian {
		t.Skip("Test only works on little-endian systems - skipping.")
	}

	defer setup(t)()

	config := getBaseConfig()
	config["login.wtmp_file_pattern"] = "../../../tests/files/wtmp"
	config["login.btmp_file_pattern"] = ""
	f := mbtest.NewReportingMetricSetV2(t, config)

	events, errs := mbtest.ReportingFetchV2(f)
	if len(errs) > 0 {
		t.Fatalf("received error: %+v", errs[0])
	}

	if len(events) == 0 {
		t.Fatal("no events were generated")
	} else if len(events) != 1 {
		t.Fatalf("only one event expected, got %d", len(events))
	}

	fullEvent := mbtest.StandardizeEvent(f, events[0], core.AddDatasetToEvent)
	mbtest.WriteEventToDataJSON(t, fullEvent, "")
}

func TestFailedLogins(t *testing.T) {
	if byteOrder != binary.LittleEndian {
		t.Skip("Test only works on little-endian systems - skipping.")
	}

	defer setup(t)()

	config := getBaseConfig()
	config["login.wtmp_file_pattern"] = ""
	config["login.btmp_file_pattern"] = "../../../tests/files/btmp_ubuntu1804"
	f := mbtest.NewReportingMetricSetV2(t, config)

	events, errs := mbtest.ReportingFetchV2(f)
	if len(errs) > 0 {
		t.Fatalf("received error: %+v", errs[0])
	}

	if len(events) == 0 {
		t.Fatal("no events were generated")
	} else if len(events) != 4 {
		t.Fatalf("expected 4 events, got %d", len(events))
	}

	// utmpdump: [6] [03307] [    ] [root    ] [ssh:notty   ] [10.0.2.2            ] [10.0.2.2       ] [2019-02-20T17:42:26,000000+0000]
	checkFieldValue(t, events[0].RootFields, "event.kind", "event")
	checkFieldValue(t, events[0].RootFields, "event.action", "user_login")
	checkFieldValue(t, events[0].RootFields, "event.outcome", "failure")
	checkFieldValue(t, events[0].RootFields, "process.pid", 3307)
	checkFieldValue(t, events[0].RootFields, "source.ip", "10.0.2.2")
	checkFieldValue(t, events[0].RootFields, "user.id", 0)
	checkFieldValue(t, events[0].RootFields, "user.name", "root")
	checkFieldValue(t, events[0].RootFields, "user.terminal", "ssh:notty")
	assert.True(t, events[0].Timestamp.Equal(time.Date(2019, 2, 20, 17, 42, 26, 0, time.UTC)),
		"Timestamp is not equal: %+v", events[0].Timestamp)

	// The second UTMP entry in the btmp test file is a duplicate of the first, this is what Ubuntu 18.04 generates.
	// utmpdump: [6] [03307] [    ] [root    ] [ssh:notty   ] [10.0.2.2            ] [10.0.2.2       ] [2019-02-20T17:42:26,000000+0000]
	checkFieldValue(t, events[1].RootFields, "event.kind", "event")
	checkFieldValue(t, events[1].RootFields, "event.action", "user_login")
	checkFieldValue(t, events[1].RootFields, "event.outcome", "failure")
	checkFieldValue(t, events[1].RootFields, "process.pid", 3307)
	checkFieldValue(t, events[1].RootFields, "source.ip", "10.0.2.2")
	checkFieldValue(t, events[1].RootFields, "user.id", 0)
	checkFieldValue(t, events[1].RootFields, "user.name", "root")
	checkFieldValue(t, events[1].RootFields, "user.terminal", "ssh:notty")
	assert.True(t, events[1].Timestamp.Equal(time.Date(2019, 2, 20, 17, 42, 26, 0, time.UTC)),
		"Timestamp is not equal: %+v", events[1].Timestamp)

	// utmpdump: [7] [03788] [/0  ] [elastic ] [pts/0       ] [                    ] [0.0.0.0        ] [2019-02-20T17:45:08,447344+0000]
	checkFieldValue(t, events[2].RootFields, "event.kind", "event")
	checkFieldValue(t, events[2].RootFields, "event.action", "user_login")
	checkFieldValue(t, events[2].RootFields, "event.outcome", "failure")
	checkFieldValue(t, events[2].RootFields, "process.pid", 3788)
	checkFieldValue(t, events[2].RootFields, "source.ip", "0.0.0.0")
	checkFieldValue(t, events[2].RootFields, "user.name", "elastic")
	checkFieldValue(t, events[2].RootFields, "user.terminal", "pts/0")
	assert.True(t, events[2].Timestamp.Equal(time.Date(2019, 2, 20, 17, 45, 8, 447344000, time.UTC)),
		"Timestamp is not equal: %+v", events[2].Timestamp)

	// utmpdump: [7] [03788] [/0  ] [UNKNOWN ] [pts/0       ] [                    ] [0.0.0.0        ] [2019-02-20T17:45:15,765318+0000]
	checkFieldValue(t, events[3].RootFields, "event.kind", "event")
	checkFieldValue(t, events[3].RootFields, "event.action", "user_login")
	checkFieldValue(t, events[3].RootFields, "event.outcome", "failure")
	checkFieldValue(t, events[3].RootFields, "process.pid", 3788)
	checkFieldValue(t, events[3].RootFields, "source.ip", "0.0.0.0")
	contains, err := events[3].RootFields.HasKey("user.id")
	if assert.NoError(t, err) {
		assert.False(t, contains)
	}
	checkFieldValue(t, events[3].RootFields, "user.name", "UNKNOWN")
	checkFieldValue(t, events[3].RootFields, "user.terminal", "pts/0")
	assert.True(t, events[3].Timestamp.Equal(time.Date(2019, 2, 20, 17, 45, 15, 765318000, time.UTC)),
		"Timestamp is not equal: %+v", events[3].Timestamp)
}

func checkFieldValue(t *testing.T, mapstr common.MapStr, fieldName string, fieldValue interface{}) {
	value, err := mapstr.GetValue(fieldName)
	if assert.NoError(t, err) {
		switch v := value.(type) {
		case *net.IP:
			assert.Equal(t, fieldValue, v.String())
		default:
			assert.Equal(t, fieldValue, v)
		}

	}
}

func getBaseConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":   "system",
		"datasets": []string{"login"},
	}
}

// setup is copied from file_integrity/metricset_test.go.
// TODO: Move to shared location and use in all unit tests.
func setup(t testing.TB) func() {
	// path.data should be set so that the DB is written to a predictable location.
	var err error
	paths.Paths.Data, err = ioutil.TempDir("", "beat-data-dir")
	if err != nil {
		t.Fatal()
	}
	return func() { os.RemoveAll(paths.Paths.Data) }
}
