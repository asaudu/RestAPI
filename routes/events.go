package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"addyCodes.com/RestAPI/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context, db *sql.DB) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	event, err := models.GetEventById(db, eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context, db *sql.DB) {
	events, err := models.GetAllEvents(db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context, db *sql.DB) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save(db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could create event"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
