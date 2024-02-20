package usecase

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

type TodoUsecase interface {
	Create(ctx context.Context, p dto.CreateTaskParams) (task *domain.Task, err error)
	Get(ctx context.Context, p dto.GetTaskParams) (task *domain.Task, err error)
	GetMany(ctx context.Context, p dto.GetTasksParams) (list *domain.TaskList, err error)
	Update(ctx context.Context, p dto.UpdateTaskParams) (task *domain.Task, err error)
	Delete(ctx context.Context, p dto.DeleteTaskParams) (err error)
}
