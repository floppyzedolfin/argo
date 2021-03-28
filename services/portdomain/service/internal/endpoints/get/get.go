package get

import (
	"context"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/db"
)

func Get(_ context.Context, _ *client.GetRequest, database db.Database) (*client.GetResponse, error) {
	logger.Log(logger.Info, "in get endpoint")
	dbPorts := database.Get()
	ports := make([]*client.Port, 0, len(dbPorts))
	for index := range dbPorts {
		p :=  dbPorts[index]
		ports = append(ports, &p)
	}
	return &client.GetResponse{Ports: ports}, nil
}
