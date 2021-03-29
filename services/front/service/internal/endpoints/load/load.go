package load

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
)

// Load calls the backend for each entry in the input file
func Load(ctx *fiber.Ctx, req Request) (Response, *fiber.Error) {
	logger.Log(logger.Info, "in load endpoint")
	f, err := os.Open(req.Input)
	if err != nil {
		return Response{}, fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("unable to open file '%s': %s", req.Input, err.Error()))
	}
	defer f.Close()
	u := upserter{backendUpserter: client.Upsert}
	if err = u.upsertPorts(ctx.Context(), f); err != nil {
		return Response{}, fiber.NewError(fiber.StatusFailedDependency, fmt.Sprintf("unable to upsert ports: %s", err.Error()))
	}

	return Response{}, nil
}
