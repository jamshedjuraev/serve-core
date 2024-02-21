package http

import (
	"errors"
	"strconv"

	"github.com/JamshedJ/backend-master-class-course/internal/delivery/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	userID := c.GetString("user_id")

	var params dto.CreateTaskParams
	err := c.BindJSON(&params)
	if err != nil {
		handleError(c, err)
		return
	}

	params.UserID = userID

	task, err := h.usecase.CreateTask(c, params)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, task)
}

func (h *Handler) getTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	userID := c.GetString("user_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.Join(errors.New("invalid value for id"), err).Error()})
	}

	task, err := h.usecase.GetTask(c, dto.GetTaskParams{
		TaskID: id,
		UserID: userID,
	})
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, task)
}

func (h *Handler) getTasks(c *gin.Context) {
	userID := c.GetString("user_id")

	var params dto.GetTasksParams
	err := c.BindQuery(&params)
	if err != nil {
		handleError(c, err)
		return
	}

	params.UserID = userID

	tasks, err := h.usecase.GetManyTasks(c, params)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, tasks)
}

func (h *Handler) updateTask(c *gin.Context) {
	idStr := c.Param("id")
	userID := c.GetString("user_id")
	var params dto.UpdateTaskParams

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, Response{Err: errors.Join(errors.New("invalid value for id"), err).Error()})
	}

	err = c.BindJSON(&params)
	if err != nil {
		handleError(c, err)
		return
	}

	params.TaskID = id
	params.UserID = userID

	task, err := h.usecase.UpdateTask(c, params)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, task)
}

func (h *Handler) deleteTask(c *gin.Context) {
	idStr := c.Param("id")
	userID := c.GetString("user_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, Response{Err: errors.Join(errors.New("invalid value for id"), err).Error()})
	}

	err = h.usecase.DeleteTask(c, dto.DeleteTaskParams{
		TaskID: id,
		UserID: userID,
	})
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "task deleted successfully"})
}
