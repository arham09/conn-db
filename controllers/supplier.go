package controllers

import (
	"fmt"
	"net/http"

	"github.com/arham09/conn-db/models"
)

type Conn interface {
	SuppliersIndex(w http.ResponseWriter, r *http.Request)
}

type Env struct {
	Db models.Datastore
}

func (env *Env) SuppliersIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	suppliers, err := env.Db.AllSuppliers()

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, supplier := range suppliers {
		fmt.Fprintf(w, "%s", supplier.Name)
	}
}
