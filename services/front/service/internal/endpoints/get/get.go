package get

import (
	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/pkg/logger"
)

func Get(ctx *fiber.Ctx, req Request) (Response,  *fiber.Error) {
	logger.Log(logger.Info, "in get endpoint")
	return Response{}, nil
}
