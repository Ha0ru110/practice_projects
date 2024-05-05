package routes

import (
	"frstapi.com/eventorganisersystem/models"
	"frstapi.com/eventorganisersystem/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse req data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse rq data"})
		return
	}
	context.JSON(201, gin.H{"message": "user created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"message": "could not handle rqst data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(401, gin.H{"message": "Either email or password invalid"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(500, gin.H{"message": "User auth failed"})
	}

	context.JSON(200, gin.H{"message": "login successful", "token": token})
}
