package controllers

import (
	"strconv"
	"zxsttm/database/models"

	"zxsttm/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProjectGetAll(context *gin.Context) {

	db := database.DB

	page, err := strconv.Atoi(context.DefaultQuery("page", "0"))

	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid page number",
		})
		return
	}

	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))

	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid page size",
		})
		return
	}

	offset := (page - 1) * pageSize

	var projects []models.Project
	db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}).Find(&projects)

	context.JSON(200, gin.H{
		"message":  "Get all projects",
		"projects": projects,
	})
}
