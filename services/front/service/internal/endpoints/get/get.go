package get

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

// Get calls the backend's Get endpoint, and cleans up "our" fields from its response
func Get(ctx *fiber.Ctx) (Response, *fiber.Error) {
	logger.Log(logger.Info, "in get endpoint")
	res, err := client.Get(ctx.Context())
	if err != nil {
		return Response{}, fiber.NewError(fiber.StatusFailedDependency, fmt.Sprintf("unable to get all ports: %s", err.Error()))
	}
	logger.Log(logger.Debug, "retrieved %d ports from the backend", len(res.Ports))
	ports := make(map[string]portdomain.Port, len(res.Ports))
	for _, p := range res.Ports {
		logger.Log(logger.Debug, "dealing with %s", p.Id)
		id := p.Id
		p.Id = ""
		ports[id] = *p
	}
	return ports, nil
}
