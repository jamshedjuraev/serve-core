package http

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	claims, err := h.userUC.ParseToken(c, token)
	if err != nil {
		c.AbortWithStatusJSON(401, Response{Err: errors.New("cannot parse token").Error()})
		return
	}
	c.Set("user_id", claims.User.ID)
}
