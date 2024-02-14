package http

import (
	"github.com/gin-gonic/gin"
)

func InitHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// user := router.Group("/auth")
	// {
	// 	user.POST("")
	// 	user.POST("")
	// }

	// task := router.Group("task")
	// {
	// 	task.POST("")
	// 	task.GET("")
	// 	task.GET("")
	// 	task.PUT("")
	// 	task.DELETE("")
	// }

	return router
}