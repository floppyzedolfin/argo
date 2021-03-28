package upsert

import (
	"context"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/db"
)

func Upsert(_ context.Context, req *client.UpsertRequest, database db.Database) (*client.UpsertResponse, error) {
	logger.Log(logger.Info, "in upsert endpoint")
	if req == nil || req.Port == nil {
		// an addition of nothing, or an upsert of nothing isn't really an error
		return nil, nil
	}
	logger.Log(logger.Info, "upserting port %s", req.Port.Id)
	database.Upsert(*req.Port)
	return nil, nil
}