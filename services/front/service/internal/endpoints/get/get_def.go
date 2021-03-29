package get

import (
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

type Request struct{}

type Response map[string]portdomain.Port
