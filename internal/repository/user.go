package repository

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"gorm.io/gorm"
)

// Check if UserRepository implements UserRepo
var _ UserRepo = (*UserRepository)(nil)

type UserRepository struct {
	db *gorm.DB
	UserRepo
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, p dto.AuthParams) (err error) {
	q := r.db.WithContext(ctx)

	var user = &domain.User{
		Username: p.Username,
		Password: p.Password,
	}

	err = q.Model(&domain.User{}).Create(&user).Error
	return
}

func (r *UserRepository) Get(ctx context.Context, p dto.AuthParams) (user *domain.User, err error) {
	q := r.db.WithContext(ctx)

	err = q.Model(&domain.User{}).Where("username = ?", p.Username).First(&user).Error
	return
}