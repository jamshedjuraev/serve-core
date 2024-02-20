package repository

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

type TodoRepo interface {
	Create(ctx context.Context, p dto.CreateTaskParams) (task *domain.Task, err error)
	Get(ctx context.Context, p dto.GetTaskParams) (task *domain.Task, err error)
	GetMany(ctx context.Context, p dto.GetTasksParams) (tasks []*domain.Task, err error)
	Update(ctx context.Context, id int, t *domain.Task) (err error)
}