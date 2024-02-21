package usecase

import (
	"context"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserRepository(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Signup(ctx context.Context, p dto.SignupParams) (err error) {
	if err = p.Validate(); err != nil {
		return err
	}
	err = u.userRepo.Signup(ctx, p)
	return
}