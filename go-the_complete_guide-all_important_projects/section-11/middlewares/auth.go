package middlewares

import (
	"frstapi.com/eventorganisersystem/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{"message": "No token found"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"message": "not authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
