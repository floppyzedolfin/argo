package db

import (
	"testing"

	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d := New()
	dCasted, ok := d.(*database)
	assert.True(t, ok)
	assert.NotNil(t, dCasted.ports)
}

func TestDatabase_Get(t *testing.T) {
	tt := map[string]struct {
		knownPorts    []string
		expectedCount int
	}{
		"no known port": {
			knownPorts:    nil,
			expectedCount: 0,
		},
		"one port": {
			knownPorts:    []string{"BFOUA"},
			expectedCount: 1,
		},
		"two ports": {
			knownPorts:    []string{"BFOUA", "GBTYN"},
			expectedCount: 2,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			d := database{ports: map[string]portdomain.Port{}}
			for _, k := range tc.knownPorts {
				d.ports[ports[k].Id] = ports[k]
			}
			res := d.Get()
			assert.Equal(t, tc.expectedCount, len(res))
		})
	}
}

func TestDatabase_Upsert(t *testing.T) {
	tt := map[string]struct {
		upserts       []string
		expectedCount int
	}{
		"nothing to upsert": {
			upserts:       nil,
			expectedCount: 0,
		},
		"one port": {
			upserts:       []string{"BFOUA"},
			expectedCount: 1,
		},
		"two ports": {
			upserts:       []string{"BFOUA", "GBTYN"},
			expectedCount: 2,
		},
		"overwrite": {
			upserts:       []string{"BFOUA", "BFOUAOverriden"},
			expectedCount: 1,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			d := database{ports: map[string]portdomain.Port{}}
			for _, k := range tc.upserts {
				d.Upsert(ports[k])
			}
			assert.Equal(t, tc.expectedCount, len(d.ports))
		})
	}
}

var ports = map[string]portdomain.Port{
	"BFOUA": {
		Name:        "Ouagadougou",
		City:        "Ouagadougou",
		Country:     "Burkina Faso",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float32{-1.5196603, 12.3714277},
		Province:    "Centre",
		Timezone:    "Africa/Ouagadougou",
		Unlocs:      []string{"BFOUA"},
		Code:        "",
		Id:          "BFOUA",
	},
	"BFOUAOverriden": {
		Name:        "Ouagadougou2",
		City:        "Ouagadougou2",
		Country:     "Burkina Faso2",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float32{-1.51966032, 12.37142772},
		Province:    "Centre",
		Timezone:    "Africa/Ouagadougou2",
		Unlocs:      []string{"BFOUA"},
		Code:        "",
		Id:          "BFOUA", // same code as BFOUA above
	},
	"GBTYN": {
		Name:        "Tyne",
		City:        "Tyne",
		Country:     "United Kingdom",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float32{-1.43, 55},
		Province:    "Newcastle upon Tyne",
		Timezone:    "Europe/London",
		Unlocs:      []string{"GBTYN"},
		Code:        "",
		Id:          "GBTYN",
	},
}
