package controllers

import (
	"zxsttm/database"
	"zxsttm/database/models"
	"zxsttm/pkg/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(context *gin.Context) {

	db := database.DB

	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var user models.User

	db.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		context.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		context.JSON(401, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Create the token with claims
	token, err := jwt.GenerateAccessToken(user)

	if err != nil {
		context.JSON(500, gin.H{
			"error": "Error while generating the token",
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "Auth Login",
		"token":   token,
	})
}
