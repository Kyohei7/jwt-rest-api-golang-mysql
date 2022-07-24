package controllers

import (
	"net/http"
	"rest-api-golang-jwt/database"
	"rest-api-golang-jwt/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {

	var user models.User

	// Check Input User -> JSON
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		context.Abort()
		return
	}

	// User Hash Password
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		context.Abort()
		return
	}

	// Insert Data User in DB
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": record.Error.Error(),
			},
		)
		context.Abort()
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"userId":   user.ID,
			"email":    user.Email,
			"username": user.Username,
		},
	)
}
