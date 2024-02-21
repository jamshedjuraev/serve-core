package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitHandler() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	user := router.Group("/auth", h.Authenticate)
	{
		user.POST("/sign-up", h.signup)
		user.POST("/sign-in", h.signin)
	}

	task := router.Group("/task")
	{
		task.POST("", h.createTask)
		task.GET("", h.getTasks)
		task.GET("/:id", h.getTaskByID)
		task.PUT("/:id", h.updateTask)
		task.DELETE("/:id", h.deleteTask)
	}
		return router
}