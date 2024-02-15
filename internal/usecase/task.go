package usecase

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
)

type TaskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewTaskInteractor(taskRepo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (t *TaskUseCase) CreateTask(ctx context.Context, params dto.TaskParams) (task *domain.Task, err error) {
	task, err = t.taskRepo.Create(ctx, params)
	return 
}