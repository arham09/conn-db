package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arham09/conn-db/config/database"
	"github.com/arham09/conn-db/config/redis"
	fr "github.com/arham09/conn-db/faktur/repository"
	"github.com/arham09/conn-db/helpers/caching"
	mid "github.com/arham09/conn-db/middleware"
	sh "github.com/arham09/conn-db/supplier/delivery/http"
	sr "github.com/arham09/conn-db/supplier/repository"
	su "github.com/arham09/conn-db/supplier/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	godotenv.Load()

	if os.Getenv(`ENV`) == `development` {
		fmt.Println("Running in development mode")
	}

}

func main() {
	dbHost := os.Getenv(`DB_HOST`)
	dbUser := os.Getenv(`DB_USER`)
	dbPassword := os.Getenv(`DB_PASSWORD`)
	dbName := os.Getenv(`DB_NAME`)

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

	redis := redis.Connect("localhost:6379", "")

	func() {
		_, err := redis.Ping(redis.Context()).Result()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	e := echo.New()

	//GlobalInterface
	cache := caching.NewRedisCaching(redis)

	// Init middleware for handler
	middl := mid.InitMiddleware(cache)

	// GlobalMiddleware
	e.Use(middleware.Gzip())
	// e.Use(middleware.Logger())

	// Repository
	supplierRepo := sr.NewPgSupplierRepository(db)
	fakturRepo := fr.NewPgFakturRepository(db)

	timeoutContext := time.Duration(2) * time.Second
	// Usecase
	supplierUsecase := su.NewSupplierUsecase(supplierRepo, fakturRepo, cache, timeoutContext)

	// Handler
	sh.NewSupplierHandler(e, supplierUsecase, middl)

	log.Fatal(e.Start(os.Getenv(`PORT`)))
}
