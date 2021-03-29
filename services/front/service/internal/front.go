package internal

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/floppyzedolfin/argo/services/front/service/internal/handlers"
)

type Service struct {
	app *fiber.App
}

// New returns a ready-to-use gateway service, that should be used with Listen()
func New() *Service {
	var g Service
	g.app = fiber.New()
	g.registerEndpoints()
	return &g
}

// Listen starts to listen to a given port
func (g *Service) Listen(port int) {
	g.app.Listen(fmt.Sprintf(":%d", port))
}

// registerEndpoints registers all the endpoints of the Gateway service
func (g *Service) registerEndpoints() {
	g.app.Post("/ports", handlers.Load)
	g.app.Get("/ports", handlers.Get)
}
