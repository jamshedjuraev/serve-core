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

func (r *TaskRepository) Create(ctx context.Context, p dto.TaskParams) (task *domain.Task, err error) {
	task = &domain.Task{
		Title:       p.Title,
		Description: p.Description,
	}

	err = r.db.WithContext(ctx).Create(&task).Error
	return
}

func (r *TaskRepository) GetMany(ctx context.Context, p dto.TaskParams) (tasks []*domain.Task, err error) {
	q := r.db.WithContext(ctx)

	if p.WithPagination {
		q = q.Offset(p.Offset()).Limit(p.PerPage)
	}

	err = q.Find(&tasks).Error
	return
}
