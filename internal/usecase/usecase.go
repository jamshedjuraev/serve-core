package usecase

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
)

type Usecase struct {
	repo repository.Repository
	TodoUsecase
	AuthUsecase
}

func NewUsecase(repo repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

type TodoUsecase interface {
	CreateTask(ctx context.Context, p dto.CreateTaskParams) (task *domain.Task, err error)
	GetTask(ctx context.Context, p dto.GetTaskParams) (task *domain.Task, err error)
	GetManyTasks(ctx context.Context, p dto.GetTasksParams) (list *domain.TaskList, err error)
	UpdateTask(ctx context.Context, p dto.UpdateTaskParams) (task *domain.Task, err error)
	DeleteTask(ctx context.Context, p dto.DeleteTaskParams) (err error)
}

type AuthUsecase interface {
	Signup(ctx context.Context, p dto.AuthParams) (err error)
	AuthenticateUser(ctx context.Context, p dto.AuthParams) (user *domain.User, err error)
	ParseToken(ctx context.Context, jwtStr string) (claims *dto.JWTClaims, err error)
}
