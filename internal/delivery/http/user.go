package http

import (
	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signup(c *gin.Context) {
	var params dto.SignupParams
	if err := c.BindJSON(&params); err != nil {
		handleError(c, err)
		return
	}

	err := h.userUC.Signup(c, params)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, gin.H{"message": "user created successfully"})
}
