// Package client exposes the API of the backend
package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

const (
	// this IP address is set in the Makefile
	ip         = "172.18.0.23"
	portNumber = 8405
)

// Get exposes the service's Get endpoint
func Get(ctx context.Context) (*portdomain.GetResponse, error) {
	clientConnexion, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, portNumber), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to create connection to portdomain: %w", err)
	}
	defer clientConnexion.Close()

	client := portdomain.NewPortdomainClient(clientConnexion)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := client.Get(ctx, &portdomain.GetRequest{})
	if err != nil {
		return nil, fmt.Errorf("error while querying the get endpoint: %w", err)
	}

	return res, nil
}

// Upsert exposes the service's Upsert endpoint
func Upsert(ctx context.Context, req *portdomain.UpsertRequest) error {
	clientConnexion, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, portNumber), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to create connection to portdomain: %w", err)
	}
	defer clientConnexion.Close()

	client := portdomain.NewPortdomainClient(clientConnexion)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err = client.Upsert(ctx, req)
	if err != nil {
		return fmt.Errorf("error while querying the upsert endpoint: %w", err)
	}

	return nil
}
