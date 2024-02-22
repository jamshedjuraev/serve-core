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
	logger := u.logger.With().Ctx(ctx).Str("method", "Signup").Logger()

	if err = p.Validate(); err != nil {
		logger.Error().Err(err).Msg("usecase.Signup p.Validate()")
		return errors.Join(ErrValidationFailed, err)
	}

	user := domain.User{
		Username: p.Username,
		Password: generatePasswordHash(p.Password),
	}

	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		logger.Error().Err(err).Msg("usecase.Signup u.repo.CreateUser()")
	}
	return
}

func (u *Usecase) AuthenticateUser(ctx context.Context, p dto.AuthParams) (user *domain.User, err error) {
	logger := u.logger.With().Ctx(ctx).Str("method", "AuthenticateUser").Logger()

	if err = p.Validate(); err != nil {
		logger.Error().Err(err).Msg("usecase.AuthenticateUser p.Validate()")
		return nil, errors.Join(ErrValidationFailed, err)
	}

	user, err = u.repo.GetUser(ctx, p)
	if err != nil {
		logger.Error().Err(err).Msg("usecase.AuthenticateUser u.repo.GetUser()")
		return nil, errors.Join(ErrInternalDatabaseError, err)
	}

	if !CheckUsersPasswordHash(user.Password, p.Password) {
		logger.Error().Err(err).Msg("usecase.AuthenticateUser CheckUsersPasswordHash()")
		return nil, errors.Join(ErrInvalidPassword, err)
	}
	return
}

func (u *Usecase) ParseToken(ctx context.Context, jwtStr string) (claims dto.JWTClaims, err error) {
	logger := u.logger.With().Ctx(ctx).Str("method", "ParseToken").Logger()

	token, err := jwt.ParseWithClaims(jwtStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if sm, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || sm != jwt.SigningMethodHS256 {
			logger.Error().Err(err).Msg("usecase.ParseToken jwt.ParseWithClaims()")
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		logger.Error().Err(err).Msg("usecase.ParseToken jwt.ParseWithClaims()")
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
