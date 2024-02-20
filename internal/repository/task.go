package repository

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"gorm.io/gorm"
)

var _ TodoRepo = (*TaskRepository)(nil)

type TaskRepository struct {
	db *gorm.DB
	TodoRepo
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, params dto.TaskParams) (task *domain.Task, err error) {
	task = &domain.Task{
		Title:       params.Title,
		Description: params.Description,
	}
	
	err = r.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		return nil, err
	}
	return
}
