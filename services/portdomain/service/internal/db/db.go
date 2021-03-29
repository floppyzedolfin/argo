package db

import (
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

// Database exposes the API we want
type Database interface {
	Upsert(portdomain.Port)
	Get() map[string]portdomain.Port
}

// database is the inner structure of our database. If we want to move to something else - mongo, ... - we can change the implemetation here.
type database struct {
	ports map[string]portdomain.Port
}

// New returns a fully functional Database
func New() Database {
	return &database{ports: make(map[string]portdomain.Port)}
}

// Upsert upserts an item in the database, using its Id as the key
func (d *database) Upsert(port portdomain.Port) {
	d.ports[port.Id] = port
}

// Get retrieves all items from the database
func (d *database) Get() map[string]portdomain.Port {
	return d.ports
}
