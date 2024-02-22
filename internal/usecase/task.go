package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

// Check if Usecase implements TodoUsecase
var _ TodoUsecase = (*Usecase)(nil)

func (u *Usecase) CreateTask(ctx context.Context, p dto.CreateTaskParams) (task *domain.Task, err error) {
	logger := u.logger.With().Ctx(ctx).Str("method", "CreateTask").Int("user_id", p.UserID).Logger()

	if err = p.Validate(); err != nil {
		logger.Error().Err(err).Msg("usecase.CreateTask p.Validate()")
		return nil, errors.Join(ErrValidationFailed, err)
	}

	task, err = u.repo.CreateTask(ctx, p)
	if err != nil {
		logger.Error().Err(err).Msg("usecase.CreateTask u.repo.CreateTask()")
		return nil, errors.Join(ErrInternalDatabaseError, err)
	}
	return
}

func (u *Usecase) GetTask(ctx context.Context, p dto.GetTaskParams) (task *domain.Task, err error) {
	logger := u.logger.With().Ctx(ctx).Str("method", "GetTask").Int("task_id", p.TaskID).Logger()
	
	if err = p.Validate(); err != nil {
		logger.Error().Err(err).Msg("usecase.GetTask p.Validate()")
		return nil, errors.Join(ErrValidationFailed, err)
	}

	task, err = u.repo.GetTask(ctx, p)
	if err != nil {
		logger.Error().Err(err).Msg("usecase.GetTask u.repo.GetTask()")
		return nil, errors.Join(ErrInternalDatabaseError, err)
	}
	return
}

func (u *Usecase) GetManyTasks(ctx context.Context, p dto.GetTasksParams) (list *domain.TaskList, err error) {
	logger := u.logger.With().Ctx(ctx).Str("method", "GetManyTasks").Int("user_id", p.UserID).Logger()

	if err = p.Validate(); err != nil {
		logger.Error().Err(err).Msg("usecase.GetManyTasks p.Validate()")
		return nil, errors.Join(ErrValidationFailed, err)
	}

	tasks, err := u.repo.GetManyTasks(ctx, p)
	if err != nil {
		logger.Error().Err(err).Msg("usecase.GetManyTasks u.repo.GetManyTasks()")
		return nil, errors.Join(ErrInternalDatabaseError, err)
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
	logger := u.logger.With().Ctx(ctx).Str("method", "UpdateTask").Int("task_id", p.TaskID).Logger()

	if err = p.Validate(); err != nil {
		logger.Error().Err(err).Msg("usecase.UpdateTask p.Validate()")
		return nil, errors.Join(ErrValidationFailed, err)
	}

	task = &domain.Task{
		Title:       p.Title,
		Description: p.Description,
		IsDone:      p.IsDone,
	}

	err = u.repo.UpdateTask(ctx, p.TaskID, task)
	if err != nil {
		logger.Error().Err(err).Msg("usecase.UpdateTask u.repo.UpdateTask()")
		return nil, errors.Join(ErrInternalDatabaseError)
	}
	return
}

func (u *Usecase) DeleteTask(ctx context.Context, p dto.DeleteTaskParams) (err error) {
	logger := u.logger.With().Ctx(ctx).Str("method", "DeleteTask").Int("task_id", p.TaskID).Logger()

	if err = p.Validate(); err != nil {
		logger.Error().Err(err).Msg("usecase.DeleteTask p.Validate()")
		return errors.Join(ErrValidationFailed, err)
	}

	isDeleted := true
	deletedAt := time.Now().UTC()

	err = u.repo.UpdateTask(ctx, p.TaskID, &domain.Task{
		IsDeleted: &isDeleted,
		DeletedAt: &deletedAt,
	})
	if err != nil {
		logger.Error().Err(err).Msg("usecase.DeleteTask u.repo.UpdateTask()")
		return errors.Join(ErrInternalDatabaseError, err)
	}
	return
}
