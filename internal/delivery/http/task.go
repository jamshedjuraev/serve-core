package http

import (
	"strconv"

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

	task, err := h.taskUC.Create(c, params); 
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, task)
}

func (h *Handler) getTasks(c *gin.Context) {
	var params dto.TaskParams
	err := c.BindQuery(&params)
	if err != nil {
		handleError(c, err)
		return
	}

	tasks, err := h.taskUC.GetMany(c, params)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, tasks)
}

func (h *Handler) getTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(c, err)
		return
	}

	task, err := h.taskUC.Get(c, dto.TaskParams{
		TaskID: id,
	})
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, task)
}