package dto

import (
	"errors"

	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

type AuthParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *AuthParams) Validate() error {
	if p.Username == "" {
		return errors.New("username is required")
	}

	if p.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type JWTClaims struct {
	jwt.RegisteredClaims
	User *domain.User
}
