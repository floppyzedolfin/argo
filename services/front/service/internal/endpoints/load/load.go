package load

import (
	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
)

func Load(ctx *fiber.Ctx, req Request) (Response, *fiber.Error) {
	logger.Log(logger.Info, "in load endpoint")
	return Response{}, nil
}
