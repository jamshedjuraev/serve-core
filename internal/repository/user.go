package repository

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Signup(ctx context.Context, p dto.SignupParams) (err error) {
	q := r.db.WithContext(ctx)

	var user = &domain.User{
		Username: p.Username,
		Password: p.Password,
	}

	err = q.Model(&domain.User{}).Create(&user).Error
	return
}