package get

import (
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

// Request is the structure of a Get request. It's empty, as the endpoint exposes the GET method
type Request struct{}

// Response is the structure of a Get response. It's similar to the contents of the input file.
type Response map[string]portdomain.Port
