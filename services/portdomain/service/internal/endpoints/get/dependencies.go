package get

import (
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

type getDatabase interface {
	Get() map[string]portdomain.Port
}
