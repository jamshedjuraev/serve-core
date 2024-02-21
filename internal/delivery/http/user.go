package http

import (
	"time"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/JamshedJ/backend-master-class-course/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) signup(c *gin.Context) {
	var params dto.AuthParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, Response{Err: err.Error()})
		return
	}

	err := h.usecase.Signup(c, params)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, gin.H{"message": "user created successfully"})
}

func (h *Handler) signin(c *gin.Context) {
	var params dto.AuthParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, Response{Err: err.Error()})
		return
	}

	u, err := h.usecase.AuthenticateUser(c, params)
	if err != nil {
		handleError(c, err)
		return
	}

	signingKey := []byte("secret")
	token, err := generateToken(c, u, signingKey)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, token)
}

func generateToken(c *gin.Context, u *domain.User, signingKey []byte) (string, error) {
	tokenTTL := 6 * time.Hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &dto.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: u,
	})
	return token.SignedString(signingKey)
}
