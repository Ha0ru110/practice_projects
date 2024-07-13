package routes

import (
	"frstapi.com/eventorganisersystem/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "event does not exist"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(500, gin.H{"message": "could not register user for event"})
		return
	}
	context.JSON(201, gin.H{"message": "Registered successfully"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(500, gin.H{"message": "could not cancel user for event"})
		return
	}
	context.JSON(200, gin.H{"message": "canceled successfully"})
}

//update
