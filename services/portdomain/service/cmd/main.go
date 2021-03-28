package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/floppyzedolfin/argo/services/portdomain/client"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/db"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/endpoints/get"
	"github.com/floppyzedolfin/argo/services/portdomain/service/internal/endpoints/upsert"
)

type portdomainServer struct {
	client.PortdomainServer
	database db.Database
}

func (p *portdomainServer) Upsert(ctx context.Context, req *client.UpsertRequest) (*client.UpsertResponse, error) {
	return upsert.Upsert(ctx, req, p.database)
}

func (p *portdomainServer) Get(ctx context.Context, req *client.GetRequest) (*client.GetResponse, error) {
	return get.Get(ctx, req, p.database)
}

func main () {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8405))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	client.RegisterPortdomainServer(grpcServer, &portdomainServer{database: db.New()})
	grpcServer.Serve(lis)
}
