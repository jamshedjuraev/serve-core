package usecase

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

// Check if UserUsecase implements AuthUsecase
var _ AuthUsecase = (*Usecase)(nil)

func (u *Usecase) Signup(ctx context.Context, p dto.AuthParams) (err error) {
	if err = p.Validate(); err != nil {
		return err
	}

	user := domain.User{
		Username: p.Username,
		Password: generatePasswordHash(p.Password),
	}

	err = u.repo.CreateUser(ctx, user)
	return
}

func (u *Usecase) AuthenticateUser(ctx context.Context, p dto.AuthParams) (user *domain.User, err error) {
	if err = p.Validate(); err != nil {
		return nil, err
	}

	user, err = u.repo.GetUser(ctx, p)
	if err != nil {
		return nil, err
	}

	if !CheckUsersPasswordHash(user.Password, p.Password) {
		return nil, errors.New("invalid password")
	}
	return
}

func (u *Usecase) ParseToken(ctx context.Context, jwtStr string) (claims dto.JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(jwtStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if sm, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || sm != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return claims, err
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
