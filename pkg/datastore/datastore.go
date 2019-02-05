package datastore

import (
	"github.com/jacobkaufmann/hallpass"
	"github.com/jmoiron/modl"
)

// A Datastore accesses the datastore (in PostgreSQL).
type Datastore struct {
	Users  hallpass.UsersService
	Passes hallpass.PassesService

	dbh modl.SqlExecutor
}

// NewDatastore creates a new client for accessing the datastore (in
// PostgreSQL). If dbh is nil, it uses the global DB handle.
func NewDatastore(dbh modl.SqlExecutor) *Datastore {
	if dbh == nil {
		dbh = DBH
	}

	d := &Datastore{dbh: dbh}
	// d.Users = &usersStore{d}
	// d.Passes = &passesStore{d}

	return d
}
