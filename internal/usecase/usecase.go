package usecase

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

type TodoUsecase interface {
	Create(ctx context.Context, params dto.TaskParams) (task *domain.Task, err error)
}
