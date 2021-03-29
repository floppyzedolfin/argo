package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/front/service/internal/endpoints/load"
)

// Load wraps the front service's load endpoint, and is in charge of decoding the request and setting the status
func Load(ctx *fiber.Ctx) error {
	req := new(load.Request)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("unable to parse body as a request: %s", err.Error())})
	}
	logger.Log(logger.Info, "received request for endpoint Load: %#v	", req)

	// call the endpoint
	res, err := load.Load(ctx, *req)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)
}
