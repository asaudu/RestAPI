package main

import (
	"net/http"

	"addyCodes.com/RestAPI/db"
	"addyCodes.com/RestAPI/models"
	"github.com/gin-gonic/gin"
)

var dbInstance = db.Database{}

var dbOperations = db.NewDatabase(dbInstance.DB)

func main() {
	dbOperations.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents(dbOperations.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save(dbOperations.DB)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could create event"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
