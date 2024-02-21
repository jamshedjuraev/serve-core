package repository

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

// Check if TaskRepository implements TodoRepo
var _ TodoRepo = (*Repository)(nil)

func (r *Repository) CreateTask(ctx context.Context, p dto.CreateTaskParams) (task *domain.Task, err error) {
	q := r.db.WithContext(ctx)

	task = &domain.Task{
		UserID:      p.UserID,
		Title:       p.Title,
		Description: p.Description,
	}

	err = q.Model(&domain.Task{}).Create(&task).Error
	return
}

func (r *Repository) GetTask(ctx context.Context, p dto.GetTaskParams) (task *domain.Task, err error) {
	q := r.db.WithContext(ctx)

	err = q.Model(&domain.Task{}).First(&task, p.TaskID).Error
	return
}

func (r *Repository) GetManyTasks(ctx context.Context, p dto.GetTasksParams) (tasks []*domain.Task, err error) {
	q := r.db.WithContext(ctx)

	if p.WithPagination {
		q = q.Offset(p.Offset()).Limit(p.PerPage)
	}

	err = q.Model(&domain.Task{}).Where("user_id = ?", p.UserID).Find(&tasks).Error
	return
}

func (r *Repository) UpdateTask(ctx context.Context, id int, t *domain.Task) (err error) {
	q := r.db.WithContext(ctx)

	err = q.Model(&domain.Task{}).Where("id = ?", id).Updates(&t).Error
	return
}
