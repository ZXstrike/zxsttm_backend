package auth

import (
	"strings"
	"zxsttm/database"
	"zxsttm/database/models"
	"zxsttm/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header["Authorization"]

		if authHeader == nil {
			context.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			context.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")

		data, err := jwt.VerifyToken(tokenString)

		if data == nil || err != nil {
			context.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			context.Abort()
			return
		}

		user := models.User{}

		if condition := database.DB.First(&user, data.UserID); condition.Error != nil {
			context.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			context.Abort()
			return

		}

		context.Set("user", user)

		context.Next()
	}
}
