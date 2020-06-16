package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arham09/conn-db/config"
	"github.com/arham09/conn-db/supplier/handler"
)

func main() {
	// var handler controllers.Conn
	db, err := config.NewDB("postgres://medea:developer@127.0.0.1/battlefield?sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	handlers := handler.Env{Db: db}
	// handler = &controllers.Env{Db: db}
	// env := &controllers.Env{Db: db}

	http.HandleFunc("/suppliers", handlers.SuppliersIndex)
	// http.HandleFunc("/suppliers", env.SuppliersIndex)
	fmt.Print("Run in 2020")
	http.ListenAndServe(":2020", nil)
}
