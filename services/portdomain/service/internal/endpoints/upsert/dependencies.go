package upsert

import (
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

type upsertDatabase interface {
	Upsert(portdomain.Port)
}
