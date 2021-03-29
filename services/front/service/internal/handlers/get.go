package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/front/service/internal/endpoints/get"
)

// Get wraps the front service's get endpoint, and is in charge of decoding the request and setting the status
func Get(ctx *fiber.Ctx) error {
	logger.Log(logger.Info, "received request for endpoint Get")

	// call the endpoint
	res, err := get.Get(ctx)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}
