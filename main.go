package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arham09/conn-db/config/database"
	fr "github.com/arham09/conn-db/faktur/repository"
	mid "github.com/arham09/conn-db/middleware"
	sh "github.com/arham09/conn-db/supplier/delivery/http"
	sr "github.com/arham09/conn-db/supplier/repository"
	su "github.com/arham09/conn-db/supplier/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`DEBUG`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`DB_HOST`)
	dbUser := viper.GetString(`DB_USER`)
	dbPassword := viper.GetString(`DB_PASSWORD`)
	dbName := viper.GetString(`DB_NAME`)

	dsn := fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`, dbUser, dbPassword, dbHost, dbName)

	db, err := database.NewDB(dsn)

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

	middl := mid.InitMiddleware()
	e.Use(middl.CORS)
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())

	// Module wiring db for repository and usecase to be used in handler
	supplierRepo := sr.NewPgSupplierRepository(db)
	fakturRepo := fr.NewPgFakturRepository(db)

	timeoutContext := time.Duration(2) * time.Second
	supplierUsecase := su.NewSupplierUsecase(supplierRepo, fakturRepo, timeoutContext)

	// Handler
	sh.NewSupplierHandler(e, supplierUsecase)

	log.Fatal(e.Start(viper.GetString(`PORT`)))
}
