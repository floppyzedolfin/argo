package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/front/service/internal/endpoints/get"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
)

func Get(ctx *fiber.Ctx) error {
	logger.Log(logger.Info, "received request for endpoint Get")

	cli := client.NewPortdomainClient(nil)
	// call the endpoint
	res, err := get.Get(ctx, get.Request{}, cli)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}
