package upsert

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/db"
)

func TestUpsert(t *testing.T) {
	tt := map[string]struct {
		reqs   []portdomain.UpsertRequest
		dbSize int
	}{
		"a single request": {
			reqs: []portdomain.UpsertRequest{
				{Port: &portdomain.Port{
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
				}}},
			dbSize: 1,
		},
		"nil request": {
			reqs:   []portdomain.UpsertRequest{},
			dbSize: 0,
		},
		"overwrite": {
			reqs: []portdomain.UpsertRequest{
				{Port: &portdomain.Port{
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
				}},
				{Port: &portdomain.Port{
					Name:        "Ouagadougou2",
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
				}},
			},
			dbSize: 1,
		},
		"two entries": {
			reqs: []portdomain.UpsertRequest{
				{Port: &portdomain.Port{
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
				}},
				{Port: &portdomain.Port{
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
				}},
			},
			dbSize: 2,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			database := db.New()
			for _, req := range tc.reqs {
				_, err := Upsert(context.Background(), &req, database)
				// obviously, this isn't returning an error, but let's just make sure it doesn't.
				assert.NoError(t, err)
			}
			// count the number of items in the DB
			assert.Equal(t, tc.dbSize, len(database.Get()))
		})
	}
}
