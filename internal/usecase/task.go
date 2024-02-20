package usecase

import (
	"context"
	"time"

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

func (u *TaskUsecase) Create(ctx context.Context, p dto.CreateTaskParams) (task *domain.Task, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	task, err = u.taskRepo.Create(ctx, p)
	return
}

func (u *TaskUsecase) Get(ctx context.Context, p dto.GetTaskParams) (task *domain.Task, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	task, err = u.taskRepo.Get(ctx, p)
	return
}

func (u *TaskUsecase) GetMany(ctx context.Context, p dto.GetTasksParams) (list *domain.TaskList, err error) {
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
		Page:  p.Page,
		Pages: pages,
		Tasks: tasks,
	}

	return
}

func (u *TaskUsecase) Update(ctx context.Context, p dto.UpdateTaskParams) (task *domain.Task, err error) {
	task = &domain.Task{
		Title:       p.Title,
		Description: p.Description,
		IsDone:      p.IsDone,
	}

	err = u.taskRepo.Update(ctx, p.TaskID, task)
	if err != nil {
		return nil, err
	}

	return
}

func (u *TaskUsecase) Delete(ctx context.Context, p dto.DeleteTaskParams) (err error) {
	isDeleted := true
	deletedAt := time.Now().UTC()

	err = u.taskRepo.Update(ctx, p.TaskID, &domain.Task{
		IsDeleted: &isDeleted,
		DeletedAt: &deletedAt,
	})
	return
}
