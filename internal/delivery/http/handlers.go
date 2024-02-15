package http

import (
	"github.com/JamshedJ/backend-master-class-course/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	taskUseCase usecase.TaskUseCase	
}

func NewHandler(taskUseCase usecase.TaskUseCase) *Handler {
	return &Handler{
		taskUseCase: taskUseCase,
	}
}

type Response struct {
	Err string `json:"error"`
}

func handleError(c *gin.Context, err error) {
	c.JSON(400, gin.H{"error": err.Error()})
}