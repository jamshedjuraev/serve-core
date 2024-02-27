package main

import (
	"os"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/http"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
	"github.com/JamshedJ/backend-master-class-course/internal/usecase"
	"github.com/JamshedJ/backend-master-class-course/pkg/glog"
	"github.com/JamshedJ/backend-master-class-course/pkg/kvstore"
	"github.com/JamshedJ/backend-master-class-course/pkg/natsClient"
	"github.com/nats-io/nats.go"
)

const (
	ServiceName = "backend-master-class"
)

func main() {
	logger := glog.NewTracingLogger()
	
	nclient := natsClient.NewNATSClient(ServiceName, os.Getenv("NATS_URL"))
	nclient.Js.AddStream(&nats.StreamConfig{
		Name:       "BACKEND_MASTER_CLASS",
		Subjects:   []string{"BACKEND_MASTER_CLASS.*"},
		MaxMsgSize: 1024 * 1024 * 20,
		Retention:  nats.WorkQueuePolicy,
	})
	
	kv := kvstore.NewNATSKVStore(os.Getenv("NATS_URL"))

	// dsn := "postgres://postgres:postgres@localhost:5433/postgres"
	dsn := kv.GetString("DATABASE_DSN")
	db, err := repository.InitDB(dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("error initializing db")
	}

	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(*repo, &logger)
	handler := http.NewHandler(*usecase)

	port := kv.GetString("SERVER_PORT")
	router := handler.InitHandler()
	if err := router.Run(port); err != nil {
		logger.Fatal().Err(err).Msg("error running server")
	}
}
