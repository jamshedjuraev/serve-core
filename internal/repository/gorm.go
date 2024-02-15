package repository

import (
	"log"
	"time"

	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
			NowFunc: func () time.Time {
				return time.Now().UTC()
			},
		},
	)
	if err != nil {
		log.Fatal("unable to connect to database")
	}

	if err = db.AutoMigrate(&domain.Task{}, &domain.User{}); err != nil {
		log.Fatal("error migrating database")
	}

	return
}

