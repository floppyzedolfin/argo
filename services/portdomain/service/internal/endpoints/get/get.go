package get

import (
	"context"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/db"
)

func Get(_ context.Context, _ *portdomain.GetRequest, database db.Database) (*portdomain.GetResponse, error) {
	logger.Log(logger.Info, "in get endpoint")
	dbPorts := database.Get()
	ports := make([]*portdomain.Port, 0, len(dbPorts))
	for index := range dbPorts {
		p :=  dbPorts[index]
		ports = append(ports, &p)
	}
	return &portdomain.GetResponse{Ports: ports}, nil
}
