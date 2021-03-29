package get

import (
	"context"
	"sort"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/db"
)

func Get(_ context.Context, _ *portdomain.GetRequest, database db.Database) (*portdomain.GetResponse, error) {
	logger.Log(logger.Info, "in get endpoint")
	dbPorts := database.Get()
	ports := make([]*portdomain.Port, 0, len(dbPorts))
	logger.Log(logger.Debug, "retrieved %d items from the DB", len(dbPorts))
	for id := range dbPorts {
		port := dbPorts[id]
		ports = append(ports, &port)
	}
	// make the output deterministic
	sort.Slice(ports, func(i, j int) bool { return ports[i].Id < ports[j].Id })
	return &portdomain.GetResponse{Ports: ports}, nil
}
