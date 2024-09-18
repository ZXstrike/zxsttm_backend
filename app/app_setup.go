package app

import (
	auth_routes "zxsttm/app/auth/routes"
	project_routes "zxsttm/app/project/routes"

	"github.com/gin-gonic/gin"
)

func InitApp(router *gin.Engine) {

	// add app routes here
	project_routes.ProjectRoutes(router)

	auth_routes.AuthRoutes(router)
}
