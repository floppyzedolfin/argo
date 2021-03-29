package db

import (
	"github.com/floppyzedolfin/argo/services/portdomain/client/portdomain"
)

// Database is the inner structure of our database. If we want to move to something else - mongo, ... - we can change the implemetation here.
type Database struct {
	ports map[string]portdomain.Port
}

// New returns a fully functional Database
func New() *Database {
	return &Database{ports: make(map[string]portdomain.Port)}
}

// Upsert upserts an item in the database, using its Id as the key
func (d *Database) Upsert(port portdomain.Port) {
	d.ports[port.Id] = port
}

// Get retrieves all items from the database
func (d *Database) Get() map[string]portdomain.Port {
	return d.ports
}
