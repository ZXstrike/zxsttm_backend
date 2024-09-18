package controllers

import "github.com/gin-gonic/gin"

func AuthCheck(context *gin.Context) {

	var token string

	if token = context.GetHeader("Authorization"); token == "" {
		context.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	data := context.Keys["user"]

	context.JSON(200, gin.H{
		"message": "Auth check",
		"token":   token,
		"user":    data,
	})
}
