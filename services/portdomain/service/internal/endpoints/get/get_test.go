package get

import (
	"context"
	"sort"
	"testing"

	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/db"
	"github.com/stretchr/testify/assert"
)

// TestGet tests the Get
func TestGet(t *testing.T) {
	tt := map[string]struct {
		knownPorts []string
	}{
		"no known port": {
			knownPorts: nil,
		},
		"one port": {
			knownPorts: []string{"BFOUA"},
		},
		"two ports":{
			knownPorts: []string{"GBTYN","BFOUA"}, // inserting them in reverse alphabetical order to check determinism
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			database := db.New()
			expected := make([]*portdomain.Port, 0, len(tc.knownPorts))
			for _, p := range tc.knownPorts {
				port := ports[p]
				database.Upsert(port)
				expected = append(expected, &port)
			}
			res, err := Get(context.Background(), nil, database)
			assert.NoError(t, err)
			// sort the slice to check determinism
			sort.Slice(expected, func(i, j int) bool {return expected[i].Id < expected[j].Id})
			assert.Equal(t, expected, res.Ports)
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
