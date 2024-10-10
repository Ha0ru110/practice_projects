package routes

import (
	"fmt"
	"frstapi.com/eventorganisersystem/models"
	"frstapi.com/eventorganisersystem/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("ShouldBindJSON" + err.Error())
		context.JSON(400, gin.H{"message": "Could not parse req data"})
		return
	}
	fmt.Println("ShouldBindJSON 1")
	token, _ := utils.GenerateToken(user.Email, user.ID)
	user.Token = token
	err = user.Save()
	if err != nil {
		fmt.Println("Save 2" + err.Error())
		context.JSON(400, gin.H{"message": "Could not parse rq data"})
		return
	}

	context.JSON(201, gin.H{"message": "user created successfully", "token": token})
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
	fmt.Print(user.Token)
	context.JSON(200, gin.H{"message": "login successful", "token": user.Token})
}
