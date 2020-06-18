package main

import (
	"log"
	"os"
	"time"

	"github.com/arham09/conn-db/config"
	"github.com/arham09/conn-db/middleware"
	sh "github.com/arham09/conn-db/supplier/delivery/http"
	sr "github.com/arham09/conn-db/supplier/repository"
	su "github.com/arham09/conn-db/supplier/usecase"
	"github.com/labstack/echo"
)

func main() {
	db, err := config.NewDB("postgres://medea:developer@127.0.0.1/battlefield?sslmode=disable")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	middl := middleware.InitMiddleware()
	e.Use(middl.CORS)

	// Module wiring db for repository and usecase to be used in handler
	supplierRepo := sr.NewPgSupplierRepository(db)
	timeoutContext := time.Duration(2) * time.Second
	supplierUsecase := su.NewSupplierUsecase(supplierRepo, timeoutContext)

	// Handler
	sh.NewSupplierHandler(e, supplierUsecase)

	log.Fatal(e.Start(":2002"))
}
