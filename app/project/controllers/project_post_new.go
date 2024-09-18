package controllers

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func ProjectPostNew(context *gin.Context) {

	db := database.DB

	var project models.Project

	if err := context.BindJSON(&project); err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	db.Create(&project)

	context.JSON(200, gin.H{
		"message": "Post new project",
		"project": project,
	})
}
