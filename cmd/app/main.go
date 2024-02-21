package main

import (
	"log"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/http"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
	"github.com/JamshedJ/backend-master-class-course/internal/usecase"
)

func main() {
	// ctx := context.Background()
	dsn := "postgres://postgres:postgres@localhost:5433/postgres"
	db, err := repository.InitDB(dsn)
	if err != nil {
		log.Fatal("error initialize DB")
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(*repo)
	handler := http.NewHandler(*usecase)
	
	router := handler.InitHandler()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("error running server on port :8080")
	}
}