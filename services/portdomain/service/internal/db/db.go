package db

import "github.com/floppyzedolfin/argo/services/portdomain/client"

type Database interface {
	Upsert(client.Port)
	Get() map[string]client.Port
}

type database struct {
	ports map[string]client.Port
}

func New() Database {
	return &database{}
}

func (d *database) Upsert(port client.Port) {
	d.ports[port.Id] = port
}

func (d *database) Get() map[string] client.Port {
	return d.ports
}
