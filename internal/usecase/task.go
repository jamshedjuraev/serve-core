package usecase

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
)

var _ TodoUsecase = (*TaskUsecase)(nil)

type TaskUsecase struct {
	taskRepo repository.TaskRepository
	TodoUsecase
}

func NewTaskUsecase(taskRepo repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		taskRepo: taskRepo,
	}
}

func (t *TaskUsecase) Create(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error) {
	task, err = t.taskRepo.Create(ctx, p)
	return
}

func (t *TaskUsecase) GetMany(ctx context.Context, p dto.TaskParams) (list *domain.TaskList, err error) {
	tasks, err := t.taskRepo.GetMany(ctx, p)
	if err != nil {
		return nil, err
	}

	var pages int
	if !p.WithPagination {
		p.Page = 1
		pages = 1
	} else {
		pages = (len(tasks) / p.PerPage) + 1
	}

	list = &domain.TaskList{
		Page: p.Page,
		Pages: pages,
		Tasks: tasks,
	}

	return
}
