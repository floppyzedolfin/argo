package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"


	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/front/service/internal/endpoints/load"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
)

func Load(ctx *fiber.Ctx) error {
	req := new(load.Request)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("unable to parse body as a request: %s", err.Error())})
	}
	logger.Log(logger.Info, "received request for endpoint Load")

	cli := client.NewPortdomainClient(nil)
	// call the endpoint
	res, err := load.Load(ctx, *req, cli)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)
}
