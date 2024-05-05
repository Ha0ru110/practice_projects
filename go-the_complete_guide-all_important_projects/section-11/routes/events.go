package routes

import (
	"fmt"
	"frstapi.com/eventorganisersystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not fetch events, try again later"})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not fetch event"})
		return
	}
	context.JSON(200, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse req data"})
		return
	}
	fmt.Printf("createEvent: %+v", event)
	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not create events, try again later "})
	}
	context.JSON(201, gin.H{"message": "Event created :)", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(500, gin.H{"message": "Could not fetch event"})
		return
	}

	if event.UserID != userId {
		context.JSON(401, gin.H{"message": "Event update not approved"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request data"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not update event"})
		return
	}
	context.JSON(200, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse event id"})
		return
	}
	userId := context.GetInt64("userId")

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not fetch the event"})
		return
	}
	if event.UserID != userId {
		context.JSON(401, gin.H{"message": "Event deletion not approved"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not delete the event"})
		return
	}
	context.JSON(200, gin.H{"message": "Event deleted successfully "})
}
