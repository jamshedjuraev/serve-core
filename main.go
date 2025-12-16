package main

import (
	"context"
	"os"

	"serve-core/config"
	"serve-core/infrastructure/repository/sqlx"

	"github.com/rs/zerolog"
)

const (
	// defaultConfigDir will be used only for local development
	// In production, the configfile will be taken from environment variable or k8s configmap
	defaultConfigDir = "./config"
)

func main() {
	ctx := context.Background()
	logger := zerolog.New(os.Stdout).With().Ctx(ctx).Timestamp().Logger()

	cfg := config.InitConfig(defaultConfigDir)

	db, err := sqlx.Connect(cfg.DB.Dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
	}
	defer db.Close()
}
