package get

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

func Get(ctx *fiber.Ctx) (Response,  *fiber.Error) {
	logger.Log(logger.Info, "in get endpoint")
	res, err := client.Get(ctx.Context())
	if err != nil {
		return Response{}, fiber.NewError(fiber.StatusFailedDependency, fmt.Sprintf("unable to get all ports: %s", err.Error()))
	}

	ports := make(map[string]portdomain.Port, len(res.Ports))
	for _, p := range res.Ports {
		p.Id = ""
		ports[p.Id] = *p
	}
	return ports, nil
}
