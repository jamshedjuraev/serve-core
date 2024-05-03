package main

import (
	"log"

	"github.com/JamshedJ/backend-master-class-course/internal/config"
	"github.com/JamshedJ/backend-master-class-course/internal/delivery/http"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
	"github.com/JamshedJ/backend-master-class-course/internal/usecase"
)

const (
	ServiceName = "backend-master-class"
)

func main() {
	cfg := config.MustLoad()

	// dsn := "postgres://user:password@localhost:5433/dbname"
	db, err := repository.InitDB(cfg.GormDSN)
	if err != nil {
		log.Fatal("error initializing db")
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(*repo)
	handler := http.NewHandler(*usecase)

	router := handler.InitHandler()
	if err := router.Run(cfg.HTTPServer.Port); err != nil {
		log.Fatal("error running server")
	}
}
