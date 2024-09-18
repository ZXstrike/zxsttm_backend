package project_routes

import (
	"zxsttm/app/project/controllers"
	"zxsttm/middleware/auth"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	projectGroup := router.Group("/projects")

	// Add routes here
	projectGroup.GET("/get-all", controllers.ProjectGetAll)

	projectGroup.POST("/post-new", auth.AuthMiddleWare(), controllers.ProjectPostNew)

	projectGroup.GET("/get-by-id", controllers.ProjectGetById)
}
