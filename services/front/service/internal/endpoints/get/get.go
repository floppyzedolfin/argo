package get

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
)

func Get(ctx *fiber.Ctx, req Request, cli client.PortdomainClient) (Response,  *fiber.Error) {
	logger.Log(logger.Info, "in get endpoint")
	res, err := cli.Get(ctx.Context(), nil)
	if err != nil {
		return Response{}, fiber.NewError(fiber.StatusFailedDependency, fmt.Sprintf("unable to get all ports: %s", err.Error()))
	}

	ports := make(map[string]client.Port, len(res.Ports))
	for _, p := range res.Ports {
		p.Id = ""
		ports[p.Id] = *p
	}
	return ports, nil
}
