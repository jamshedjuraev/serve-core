package usecase

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
)

// Check if TaskUsecase implements TodoUsecase
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

func (u *TaskUsecase) Create(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error) {
	task, err = u.taskRepo.Create(ctx, p)
	return
}

func (u *TaskUsecase) GetMany(ctx context.Context, p dto.TaskParams) (list *domain.TaskList, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	tasks, err := u.taskRepo.GetMany(ctx, p)
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

func (u *TaskUsecase) Get(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	task, err = u.taskRepo.Get(ctx, p)
	return
}

func (u *TaskUsecase) Update(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error) {
	task, err = u.taskRepo.Update(ctx, p)
	return
}