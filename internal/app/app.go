package app

import (
	"github.com/gin-contrib/cors"
	"log"
	"os"
	"os/signal"
	"pahan/config"
	v1 "pahan/internal/controller/http/v1"
	"pahan/internal/usecase"
	"pahan/internal/usecase/repo"
	"pahan/pkg/httpserver"
	"pahan/pkg/postgres"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	pg, err := postgres.New(cfg)
	if err != nil {
		log.Fatal("Error in creating Postgres Instance")
	}

	modelUseCase := usecase.NewModelUseCase(repo.NewModelRepo(pg))
	ordersUseCase := usecase.NewOrdersUseCase(repo.NewOrdersRepo(pg))
	shipmentUseCase := usecase.NewShipmentUseCase(repo.NewShipmentRepo(pg))
	subsidyUseCase := usecase.NewSubsidyUseCase(repo.NewSubsidyRepo(pg))
	engineerUseCase := usecase.NewEngineerUseCase(repo.NewEngineerRepo(pg))
	factoryUseCase := usecase.NewFactoryUseCase(repo.NewFactoryRepo(pg))
	componentUseCase := usecase.NewComponentUseCase(repo.NewComponentRepo(pg))
	typeUseCase := usecase.NewTypeUseCase(repo.NewTypeRepo(pg))
	loginUseCase := usecase.NewLoginUseCase(
		repo.NewVendorRepo(pg),
		repo.NewCountryRepo(pg))
	// http Server
	handler := gin.New()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v1.NewRouter(handler,
		modelUseCase,
		ordersUseCase,
		shipmentUseCase,
		subsidyUseCase,
		engineerUseCase,
		factoryUseCase,
		componentUseCase,
		typeUseCase,
		loginUseCase)

	serv := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
