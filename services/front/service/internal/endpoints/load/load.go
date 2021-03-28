package load

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

func Load(ctx *fiber.Ctx, req Request) (Response, *fiber.Error) {
	logger.Log(logger.Info, "in load endpoint")
	for id, port := range req {
		p := port
		p.Id = id

		err := client.Upsert(ctx.Context(), &portdomain.UpsertRequest{Port: &p})
		if err != nil {
			return Response{}, fiber.NewError(fiber.StatusFailedDependency, fmt.Sprintf("unable to upsert port %s: %s", id, err.Error()))
		}
	}
	return Response{}, nil
}
