package db

import (
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

type Database interface {
	Upsert(portdomain.Port)
	Get() map[string]portdomain.Port
}

type database struct {
	ports map[string]portdomain.Port
}

func New() Database {
	return &database{}
}

func (d *database) Upsert(port portdomain.Port) {
	d.ports[port.Id] = port
}

func (d *database) Get() map[string]portdomain.Port {
	return d.ports
}
