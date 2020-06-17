package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/arham09/conn-db/supplier/repository"
)

type Env struct {
	Db *sql.DB
}

func (env *Env) SuppliersIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	suppliers, err := repository.AllSuppliers(env.Db)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, supplier := range suppliers {
		fmt.Fprintf(w, "%s", supplier.Name)
	}
}
