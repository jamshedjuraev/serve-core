package usecase

import (
	"context"
	"time"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

// Check if Usecase implements TodoUsecase
var _ TodoUsecase = (*Usecase)(nil)

func (u *Usecase) CreateTask(ctx context.Context, p dto.CreateTaskParams) (task *domain.Task, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	task, err = u.repo.CreateTask(ctx, p)
	return
}

func (u *Usecase) GetTask(ctx context.Context, p dto.GetTaskParams) (task *domain.Task, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	task, err = u.repo.GetTask(ctx, p)
	return
}

func (u *Usecase) GetManyTasks(ctx context.Context, p dto.GetTasksParams) (list *domain.TaskList, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	tasks, err := u.repo.GetManyTasks(ctx, p)
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

func (u *Usecase) UpdateTask(ctx context.Context, p dto.UpdateTaskParams) (task *domain.Task, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	task = &domain.Task{
		Title:       p.Title,
		Description: p.Description,
		IsDone:      p.IsDone,
	}

	err = u.repo.UpdateTask(ctx, p.TaskID, task)
	if err != nil {
		return nil, err
	}

	return
}

func (u *Usecase) DeleteTask(ctx context.Context, p dto.DeleteTaskParams) (err error) {
	if err = p.Validate(); err != nil {
		return err
	}

	isDeleted := true
	deletedAt := time.Now().UTC()

	err = u.repo.UpdateTask(ctx, p.TaskID, &domain.Task{
		IsDeleted: &isDeleted,
		DeletedAt: &deletedAt,
	})
	return
}
