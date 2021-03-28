package upsert

import (
	"testing"

	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

func TestUpsert(t *testing.T) {
	tt := map[string]struct {
		reqs []portdomain.UpsertRequest
	} {
		"a single request": {
			reqs: []portdomain.UpsertRequest{{Port: &portdomain.Port{
				Name:        "Loyalty",
				City:        "If it be banish'd from the frosty head,",
				Country:     "Where shall it find a harbour in the earth?",
				Alias:       []string{"LYLY"},
				Regions:     nil,
				Coordinates: nil,
				Province:    "",
				Timezone:    "",
				Unlocs:      nil,
				Code:        "",
				Id:          "",
			}}},
		},
	}
}
