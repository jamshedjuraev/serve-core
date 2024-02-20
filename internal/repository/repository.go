package repository

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

type TodoRepo interface {
	Create(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error)
	GetMany(ctx context.Context, p dto.TaskParams) (tasks []*domain.Task, err error)
	Get(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error)
}