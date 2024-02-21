package repository

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
)

// Check if UserRepo implements UserRepo
var _ UserRepo = (*Repository)(nil)

func (r *Repository) CreateUser(ctx context.Context, u domain.User) (err error) {
	q := r.db.WithContext(ctx)

	err = q.Model(&domain.User{}).Create(&u).Error
	return
}

func (r *Repository) GetUser(ctx context.Context, p dto.AuthParams) (user *domain.User, err error) {
	q := r.db.WithContext(ctx)

	err = q.Model(&domain.User{}).Where("username = ?", p.Username).First(&user).Error
	return
}
