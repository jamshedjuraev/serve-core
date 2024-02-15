package http

import (
	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	var params dto.TaskParams
	err := c.BindJSON(&params)
	if err != nil {
		handleError(c, err)
		return
	}

	task, err := h.taskUseCase.CreateTask(c, params); 
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, task)
}