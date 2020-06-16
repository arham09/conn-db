package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Datastore as interface for function to get data from db
type Datastore interface {
	AllSuppliers() ([]*Supplier, error)
}

type DB struct {
	*sql.DB
}

// NewDB func to connect to DB
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
