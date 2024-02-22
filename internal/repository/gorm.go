package repository

import (
	"time"

	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/JamshedJ/backend-master-class-course/pkg/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (db *gorm.DB, err error) {
	logger := glog.NewTracingLogger()

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	},
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening database connection")
	}

	if err = db.AutoMigrate(&domain.Task{}, &domain.User{}); err != nil {
		logger.Fatal().Err(err).Msg("error migrating database schema")
	}
	return
}
