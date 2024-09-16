package routes

import (
	"net/http"
	"strconv"

	"example.com/main.go/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Issue while fetching Try again"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //10 decimal system //64 bit size
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot able to parse event ID"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "not able to fetch the event"})
		return
	}
	context.JSON(http.StatusOK, event)

}
func createEvent(context *gin.Context) {
  token := context.Request.Header.Get("Authorization")

  if token == ""{
	context.JSON(http.StatusUnauthorized, gin.H{"message": "Token missing"})
		return
  }
  

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing required fields data"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Eventcreated", "event": event})
}

func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //10 decimal system //64 bit size
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot able to parse event ID"})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing required fields data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not update event"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"Successfully updated"})
     
}

func deleteEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) //10 decimal system //64 bit size
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot able to parse event ID"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event"})
		return
	}

	err = event.Delete()

	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete the event"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Event deleted Successfully"})
}