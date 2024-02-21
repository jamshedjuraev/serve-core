package repository

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	TodoRepo
	UserRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

type TodoRepo interface {
	CreateTask(ctx context.Context, p dto.CreateTaskParams) (task *domain.Task, err error)
	GetTask(ctx context.Context, p dto.GetTaskParams) (task *domain.Task, err error)
	GetManyTasks(ctx context.Context, p dto.GetTasksParams) (tasks []*domain.Task, err error)
	UpdateTask(ctx context.Context, id int, t *domain.Task) (err error)
}

type UserRepo interface {
	CreateUser(ctx context.Context, p dto.AuthParams) (err error)
	GetUser(ctx context.Context, p dto.AuthParams) (user *domain.User, err error)
}
