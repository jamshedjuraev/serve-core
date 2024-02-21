package usecase

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/JamshedJ/backend-master-class-course/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

// Check if UserUsecase implements AuthUsecase
var _ AuthUsecase = (*UserUsecase)(nil)

type UserUsecase struct {
	userRepo repository.UserRepository
	AuthUsecase
}

func NewUserRepository(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Signup(ctx context.Context, p dto.AuthParams) (err error) {
	if err = p.Validate(); err != nil {
		return err
	}
	err = u.userRepo.Create(ctx, p)
	return
}

func (u *UserUsecase) AuthenticateUser(ctx context.Context, p dto.AuthParams) (user *domain.User, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	user, err = u.userRepo.Get(ctx, p)
	if err != nil {
		return nil, err
	}

	if !CheckUsersPasswordHash(user.Password, p.Password) {
		return nil, errors.New("invalid password")
	}
	return
}

func (u *UserUsecase) ParseToken(ctx context.Context, jwtStr string) (claims *dto.JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(jwtStr, &dto.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return
}

func CheckUsersPasswordHash(dbPass, rowPass string) bool {
	hashedPass := generatePasswordHash(rowPass)
	return dbPass == hashedPass
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte("salt")))
}
