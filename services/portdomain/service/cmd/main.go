package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/db"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/endpoints/get"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/endpoints/upsert"
)

type portdomainServer struct {
	portdomain.PortdomainServer
	database db.Database
}

func (p *portdomainServer) Upsert(ctx context.Context, req *portdomain.UpsertRequest) (*portdomain.UpsertResponse, error) {
	return upsert.Upsert(ctx, req, p.database)
}

func (p *portdomainServer) Get(ctx context.Context, req *portdomain.GetRequest) (*portdomain.GetResponse, error) {
	return get.Get(ctx, req, p.database)
}

func main () {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8405))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	portdomain.RegisterPortdomainServer(grpcServer, &portdomainServer{database: db.New()})
	logger.Log(logger.Info,"portdomain running...")
	grpcServer.Serve(lis)
}
