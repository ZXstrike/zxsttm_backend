package auth_routes

import (
	"zxsttm/app/auth/controllers"
	"zxsttm/middleware/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	authGroup := router.Group("/auth")

	// Add routes here
	authGroup.POST("/login", controllers.AuthLogin)

	authGroup.POST("/register", controllers.AuthRegister)

	authGroup.GET("/check", auth.AuthMiddleWare(), controllers.AuthCheck)

}
