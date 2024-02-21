package http

import (
	"github.com/JamshedJ/backend-master-class-course/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	taskUC usecase.TaskUsecase	
	userUC usecase.UserUsecase
}

func NewHandler(taskUC usecase.TaskUsecase) *Handler {
	return &Handler{
		taskUC: taskUC,
	}
}

type Response struct {
	Err string `json:"error"`
}

func handleError(c *gin.Context, err error) {
	c.JSON(400, gin.H{"error": err.Error()})
}