package main

import (
	"github.com/JamshedJ/backend-master-class-course/internal/delivery/http"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
	"github.com/JamshedJ/backend-master-class-course/internal/usecase"
	"github.com/JamshedJ/backend-master-class-course/pkg/glog"
)

func main() {
	logger := glog.NewTracingLogger()

	dsn := "postgres://postgres:postgres@localhost:5433/postgres"
	db, err := repository.InitDB(dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("error initializing db")
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(*repo, &logger)
	handler := http.NewHandler(*usecase)

	router := handler.InitHandler()
	if err := router.Run(":8080"); err != nil {
		logger.Fatal().Err(err).Msg("error running server")
	}
}
