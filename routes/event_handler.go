package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/go_app/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	event, err := models.GetEventData(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	user_id := context.GetInt64("user_id")
	event.UserID = user_id
	fmt.Println(event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the event",
		})
		return
	}
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could Not save Event",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	user_id := context.GetInt64("user_id")

	event, err := models.GetEventData(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if event.UserID != user_id {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update event",
		})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	updatedEvent.ID = id
	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Updating Event Failed",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	user_id := context.GetInt64("user_id")

	event, err := models.GetEventData(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Event Id",
		})
		return
	}

	if event.UserID != user_id {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to delete event",
		})
		return
	}
	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delering Event Failed",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event Deleted Successfully",
	})
}
