package main

import (
	"log"
	"os"
	"time"

	"github.com/arham09/conn-db/config"
	"github.com/arham09/conn-db/middleware"
	handler "github.com/arham09/conn-db/supplier/delivery/http"
	"github.com/arham09/conn-db/supplier/repository"
	"github.com/arham09/conn-db/supplier/usecase"
	"github.com/labstack/echo"
)

func main() {
	// var handler controllers.Conn
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

	supplierRepo := repository.NewPgSupplierRepository(db)

	timeoutContext := time.Duration(2) * time.Second

	supplierUsecase := usecase.NewSupplierUsecase(supplierRepo, timeoutContext)

	handler.NewSupplierHandler(e, supplierUsecase)

	log.Fatal(e.Start(":2002"))
}
