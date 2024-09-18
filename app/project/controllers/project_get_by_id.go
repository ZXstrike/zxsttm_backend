package controllers

import (
	"zxsttm/database"
	"zxsttm/database/models"

	"github.com/gin-gonic/gin"
)

func ProjectGetById(context *gin.Context) {

	db := database.DB

	project := models.Project{}

	db.First(&project, context.Query("id"))

	context.JSON(200, gin.H{
		"message": "Get project by ID",
		"project": project,
	})
}
