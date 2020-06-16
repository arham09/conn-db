package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arham09/conn-db/controllers"
	"github.com/arham09/conn-db/models"
)

func main() {
	var handler controllers.Conn
	db, err := models.NewDB("postgres://medea:developer@127.0.0.1/battlefield?sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	handler = &controllers.Env{Db: db}
	// env := &controllers.Env{Db: db}

	http.HandleFunc("/suppliers", handler.SuppliersIndex)
	// http.HandleFunc("/suppliers", env.SuppliersIndex)
	fmt.Print("Run in 2020")
	http.ListenAndServe(":2020", nil)
}
