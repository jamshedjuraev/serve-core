package usecase

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

type TodoUsecase interface {
	Create(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error)
	GetMany(ctx context.Context, p dto.TaskParams) (list *domain.TaskList, err error)
	Get(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error)
	Update(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error)
}
